package scm

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/mod/semver"
)

type GitError struct {
	msg   GitErrorMsg
	cause error
}

func (e GitError) Error() string {
	return fmt.Sprintf("%s - %s", string(e.msg), e.cause.Error())
}

func (e GitError) Unwrap() error {
	return e.cause
}

type GitErrorMsg string

func (e GitErrorMsg) Error() string {
	return string(e)
}

func (e GitErrorMsg) Wrap(cause error) GitError {
	return GitError{
		msg:   e,
		cause: cause,
	}
}

const (
	CheckoutFailed     GitErrorMsg = "failed to checkout the requested worktree"
	CloneFailed        GitErrorMsg = "failed to clone the repository"
	GetTagsFailed      GitErrorMsg = "failed to retrieve repository tags"
	NoReleaseTag       GitErrorMsg = "repository has no release tags matching semver"
	SourceReadFailed   GitErrorMsg = "failed to read source file"
	WorktreeWalkFailed GitErrorMsg = "error walking the worktree"
)

type Git struct {
	url        string
	repository *git.Repository
}

// Github clones remote specified by the provided organization and
// repository names as "origin" without checking out a working tree (for
// performance reasons).
func Github(organization, repository string) (*Git, error) {
	sb := strings.Builder{}
	sb.WriteString("https://github.com/")
	sb.WriteString(organization)
	sb.WriteString("/")
	sb.WriteString(repository)

	return Repository(sb.String())
}

// Repository clones remote specified by the provided URL as "origin"
// without checking out a working tree (for performance reasons).
func Repository(url string) (*Git, error) {
	storage := memory.NewStorage()
	filesystem := memfs.New()

	repo, err := git.Clone(storage, filesystem, &git.CloneOptions{
		URL:               url,
		Auth:              nil,
		RemoteName:        "origin",
		ReferenceName:     "",
		SingleBranch:      true,
		NoCheckout:        true,
		Depth:             1,
		RecurseSubmodules: git.NoRecurseSubmodules,
		Progress:          nil,
		Tags:              git.AllTags,
	})
	if err != nil {
		return nil, CloneFailed.Wrap(err)
	}

	return &Git{
		url:        url,
		repository: repo,
	}, nil
}

// Checkout retrieves the working tree for the specified reference from
// the specified repository.
func (g *Git) Checkout(reference *plumbing.Reference) (*git.Worktree, error) {
	worktree, err := g.repository.Worktree()
	if err != nil {
		return nil, CheckoutFailed.Wrap(err)
	}

	return worktree, worktree.Checkout(&git.CheckoutOptions{
		Hash:   reference.Hash(),
		Branch: "",
		Create: false,
		Force:  false,
		Keep:   false,
	})
}

// Filepaths walks the worktree's filesystem hierarchy and returns a
// slice of FileInfo values for all the elements encountered.
func (g *Git) Filepaths() ([]string, error) {
	paths := []string{}

	return paths, g.Walk(func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		paths = append(paths, path)

		return nil
	})
}

// Latest returns the reference to the latest tag matching the semver
// criteria without a prerelease suffix (highest, major, minor, bug
// values ignoring the build number suffix if present).
func (g *Git) Latest() (*plumbing.Reference, error) {
	releases, err := g.Releases()
	if err != nil {
		return nil, err
	}

	if len(releases) == 0 {
		return nil, NoReleaseTag
	}

	less := func(i, j int) bool {
		return semver.Compare(releases[i].Name().Short(), releases[j].Name().Short()) < 1
	}

	sort.SliceStable(releases, less)

	return releases[len(releases)-1], nil
}

// Releases returns the list of tags that match the semantic versioning
// criteria and do not include a prerelease suffix.
func (g *Git) Releases() ([]*plumbing.Reference, error) {
	tags, err := g.repository.Tags()
	if err != nil {
		return nil, GetTagsFailed.Wrap(err)
	}

	releases := []*plumbing.Reference{}
	_ = tags.ForEach(func(ref *plumbing.Reference) error {
		if semver.IsValid(ref.Name().Short()) && semver.Prerelease(ref.Name().Short()) == "" {
			releases = append(releases, ref)
		}

		return nil
	})

	return releases, nil
}

// Source retrieves the contents of the file at the provided path from
// the repository's current worktree.
func (g *Git) Source(path string) ([]byte, error) {
	worktree, err := g.repository.Worktree()
	if err != nil {
		return nil, SourceReadFailed.Wrap(err)
	}

	fileInfo, err := worktree.Filesystem.Stat(path)
	if err != nil {
		return nil, SourceReadFailed.Wrap(err)
	}

	file, err := worktree.Filesystem.Open(path)
	if err != nil {
		return nil, SourceReadFailed.Wrap(err)
	}
	defer file.Close()

	source := make([]byte, fileInfo.Size())
	if _, err = file.Read(source); err != nil {
		return nil, SourceReadFailed.Wrap(err)
	}

	return source, nil
}

// Walk traverses the in-memory working tree filesystem hierarchy, while
// applying the provided WalkFunc to each inspected element.
func (g *Git) Walk(fn filepath.WalkFunc) error {
	worktree, err := g.repository.Worktree()
	if err != nil {
		return WorktreeWalkFailed.Wrap(err)
	}

	billy := worktree.Filesystem

	return walk(billy, billy.Root(), fn)
}

func walk(billy billy.Filesystem, path string, fn filepath.WalkFunc) error {
	stat, err := billy.Stat(path)
	if err != nil {
		return WorktreeWalkFailed.Wrap(err)
	}

	err = fn(path, stat, err)
	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return nil
	}

	dir, err := billy.ReadDir(path)
	if err != nil {
		return WorktreeWalkFailed.Wrap(err)
	}

	for _, info := range dir {
		err = walk(billy, filepath.Join(path, info.Name()), fn)
		if err != nil {
			return err
		}
	}

	return nil
}

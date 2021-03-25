package scm_test

import (
	"strings"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/selesy/go-material/internal/scm"
	"github.com/stretchr/testify/require"
)

// nolint:tparallel,paralleltest
func TestGit(t *testing.T) {
	t.Parallel()

	var repo *scm.Git

	t.Run("Retrieve Github repository without checking out a working tree", func(t *testing.T) {
		var err error

		repo, err = scm.Github("material-components", "material-components-web")
		require.NoError(t, err)
		require.NotNil(t, repo)

		// worktree, err := repo.Worktree()
		// require.NoError(t, err)

		_, err = repo.Filepaths()
		require.EqualError(t, err, "error walking the worktree - file does not exist")
	})

	var latest *plumbing.Reference

	t.Run("Get a reference to the latest release", func(t *testing.T) {
		var err error

		latest, err = repo.Latest()
		require.NoError(t, err)
		require.Equal(t, "v10.0.0", latest.Name().Short())
	})

	t.Run("Checkout the worktree for the latest release", func(t *testing.T) {
		worktree, err := repo.Checkout(latest)
		require.NoError(t, err)
		require.NotNil(t, worktree)
	})

	t.Run("Retrieve the paths for all files in the worktree", func(t *testing.T) {
		paths, err := repo.Filepaths()
		require.NoError(t, err)
		require.NotEmpty(t, paths)
		for _, path := range paths {
			t.Log("File path: ", path)
		}
	})

	t.Run("Retrieve the source for a specific file", func(t *testing.T) {
		source, err := repo.Source("/packages/mdc-chips/README.md")
		require.NoError(t, err)
		t.Log("Source: ", string(source))
		require.True(t, strings.HasPrefix(string(source), "<!--docs:\ntitle: \"Chips\""))
		require.Contains(t, string(source), "# Chips")
	})
}

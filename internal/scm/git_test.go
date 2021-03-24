package scm_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/selesy/go-material/internal/scm"
	"github.com/stretchr/testify/require"
)

// nolint:tparallel,paralleltest
func TestGit(t *testing.T) {
	t.Parallel()

	var repo *git.Repository

	t.Run("Retrieve Github repository without checking out a working tree", func(t *testing.T) {
		var err error

		repo, err = scm.Github("material-components", "material-components-web")
		require.NoError(t, err)
		require.NotNil(t, repo)

		worktree, err := repo.Worktree()
		require.NoError(t, err)

		_, err = scm.TreeEntries(worktree)
		require.EqualError(t, err, "error walking the worktree - file does not exist")
	})

	var latest *plumbing.Reference

	t.Run("Get a reference to the latest release", func(t *testing.T) {
		var err error

		latest, err = scm.Latest(repo)
		require.NoError(t, err)
		require.Equal(t, "v10.0.0", latest.Name().Short())
	})

	var worktree *git.Worktree

	t.Run("Checkout the worktree for the latest release", func(t *testing.T) {
		var err error

		worktree, err = scm.Checkout(repo, latest)
		require.NoError(t, err)
		require.NotNil(t, worktree)
	})

	t.Run("Retrieve the information for all files in the worktree", func(t *testing.T) {
		fis, err := scm.TreeEntries(worktree)
		require.NoError(t, err)
		require.NotEmpty(t, fis)
	})
}

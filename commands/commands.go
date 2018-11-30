package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

var (
	rootFolder string
)

func addRootFolderFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&rootFolder, "root-folder", "r", gitRoot(), "folder containing the monorepo")
}

func gitRoot() string {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return ""
	}
	t, err := r.Worktree()
	if err != nil {
		return ""
	}
	return t.Filesystem.Root()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

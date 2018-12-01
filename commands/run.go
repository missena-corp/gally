package commands

import (
	"fmt"

	"github.com/missena-corp/gally/config"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "let Gally run your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := config.FindProjects(gitRoot())
		fmt.Printf("Gally monorepo handler %v", f)
	},
}

func init() {
	addRootFolderFlag(runCmd)
	rootCmd.AddCommand(runCmd)
}

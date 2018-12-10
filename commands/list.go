package commands

import (
	"fmt"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "display all handled projects",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		projects := project.FindProjects(rootDir)
		for _, p := range projects {
			fmt.Printf("* %s: %s", p.Name, p.Dir)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

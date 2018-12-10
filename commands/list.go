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
		for _, p := range project.FindProjects(rootDir) {
			fmt.Printf("* %s: %s", p.Name, p.Dir)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
)

const padding = 1

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "display all handled projects",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
		fmt.Fprintf(w, "Name\t Directory\t\n")
		for _, p := range project.FindAll(rootDir) {
			fmt.Fprintf(w, "%s\t %s\t\n", p.Name, p.BaseDir)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

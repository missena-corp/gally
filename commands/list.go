package commands

import (
	"encoding/json"
	"fmt"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
)

const padding = 1

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "display all handled projects",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		out, _ := json.MarshalIndent(project.FindAll(rootDir).ToSlice(), "", "\t")
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

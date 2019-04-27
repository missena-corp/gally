package commands

import (
	"encoding/json"
	"fmt"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

const padding = 1

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "display all handled projects",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		projects := project.Projects{}
		if projectName != "" {
			p := project.FindByName(rootDir, projectName)
			if p == nil {
				jww.ERROR.Fatalf("could not find project %q", projectName)
			}
			projects[projectName] = project.FindByName(rootDir, projectName)
		} else {
			projects = project.FindAll(rootDir)
		}
		out, _ := json.MarshalIndent(projects.ToSlice(), "", "\t")
		fmt.Println(string(out))
	},
}

func init() {
	addProjectFlag(listCmd)
	rootCmd.AddCommand(listCmd)
}

package commands

import (
	"encoding/json"
	"fmt"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var commandExistCmd = &cobra.Command{
	Use:   "script-exists [script]",
	Short: "tells is script is defined in the project",
	Run: func(cmd *cobra.Command, args []string) {
		projects := project.Projects{}
		if projectName != "" {
			p := project.FindByName(rootDir, projectName)
			if p == nil {
				jww.ERROR.Fatalf("could not find project %q", projectName)
			}
			projects[projectName] = project.FindByName(rootDir, projectName)
		} else if updated {
			projects = project.FindAllUpdated(rootDir)
		} else {
			projects = project.FindAll(rootDir)
		}
		out, _ := json.MarshalIndent(projects.ToSlice(), "", "\t")
		fmt.Println(string(out))
	},
}

func init() {
	addProjectFlag(listCmd)
	rootCmd.AddCommand(commandExistCmd)
}

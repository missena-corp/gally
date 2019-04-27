package commands

import (
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run your script on projects having updated files",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		if len(args) == 0 {
			jww.ERROR.Fatalf("no script provided in command")
		}
		var projects project.Projects
		if projectName != "" {
			p := project.FindByName(rootDir, projectName)
			if p == nil {
				jww.ERROR.Fatalf("could not find project %q", projectName)
			}
			projects[projectName] = project.FindByName(rootDir, projectName)
		} else {
			projects = project.FindAllUpdated(rootDir)
		}
		script := args[0]
		for _, p := range projects {
			if err := p.Run(script); err != nil {
				jww.ERROR.Fatalf("could not run properly script %q: %v", script, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

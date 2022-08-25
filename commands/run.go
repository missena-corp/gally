package commands

import (
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var runCmd = &cobra.Command{
	Use:     "run [script]",
	Aliases: []string{"exec"},
	Short:   "run your script on projects having updated files",
	Long: `run your script on projects having updated files.
	Be careful using without option --ignore-missing would mean the script is defined on *every* project`,
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		if len(args) == 0 {
			jww.ERROR.Fatalf("no script provided in command")
		}
		projects := project.Projects{}
		if projectName != "" {
			projects[projectName] = project.FindByName(rootDir, projectName)
			if projects[projectName] == nil {
				jww.ERROR.Fatalf("could not find project %q", projectName)
			}
		} else if allProjects {
			projects = project.FindAll(rootDir)
		} else {
			projects = project.FindAllUpdated(rootDir, noDependency)
		}
		script := args[0]
		for _, p := range projects {
			err := p.Run(script)
			if err != nil && (!ignoreMissing || err != project.ErrCmdDoesNotExist) {
				jww.ERROR.Fatalf("could not run properly script %q for %q: %v", script, p.Name, err)
			}
		}
	},
}

func init() {
	addAllProjectsFlag(runCmd)
	addIgnoreMissingFlag(runCmd)
	addNoDependencyFlag(runCmd)
	addProjectFlag(runCmd)
	rootCmd.AddCommand(runCmd)
}

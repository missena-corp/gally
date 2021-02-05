package commands

import (
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "build your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		jww.WARN.Println("using 'gally build' is deprecated now use 'gally run build'")
		handleVerboseFlag()
		if tag != "" {
			if err := project.BuildTag(projectName, tag, rootDir); err != nil {
				jww.ERROR.Fatalf("could not build properly project: %v", err)
			}
			return
		}
		jww.INFO.Println("building no tag projects")
		if err := project.BuildWithoutTag(projectName, rootDir, noDependency); err != nil {
			jww.ERROR.Fatalf("could not build no-tag projects: %v", err)
		}
		if force {
			jww.INFO.Println("building projects with tag in force mode")
			if err := project.BuildForceWithoutTag(projectName, rootDir, noDependency); err != nil {
				jww.ERROR.Fatalf("could not build tag projects: %v", err)
			}
		}
	},
}

func init() {
	addForceFlag(buildCmd)
	addNoDependencyFlag(buildCmd)
	addProjectFlag(buildCmd)
	addTagFlag(buildCmd)
	addVersionFlag(buildCmd)
	rootCmd.AddCommand(buildCmd)
}

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
		handleVerboseFlag()
		var p *string = nil
		if projectName != "" {
			p = &projectName
		}
		if tag != "" {
			if err := project.BuildTag(p, tag, rootDir); err != nil {
				jww.ERROR.Fatalf("could not build properly project %s@%s in %q: %v", *p, tag, rootDir, err)
			}
			return
		}
		jww.INFO.Println("building no tag projects")
		if err := project.BuildWithoutTag(p, rootDir, noDependency); err != nil {
			jww.ERROR.Fatalf("could not build no-tag projects: %v", err)
		}
		if force {
			jww.INFO.Println("building projects with tag in force mode")
			if err := project.BuildForceWithoutTag(p, rootDir, noDependency); err != nil {
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
	rootCmd.AddCommand(buildCmd)
}

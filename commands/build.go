package commands

import (
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
        var p *string = nil
        if projectName != "" {
            p = &projectName
        }
		if tag == "" {
			jww.INFO.Printf("building no tag projects")
			if err := project.BuildNoTag(p, rootDir); err != nil {
				jww.ERROR.Fatalf("could not build no-tag projects: %v", err)
			}
			if force {
				jww.INFO.Printf("building projects with tag in force mode")
					if err := project.BuildForceWithTag(p, rootDir); err != nil {
						jww.ERROR.Fatalf("could not build tag projects: %v", err)
					}
				}
			return
		}
		if err := project.BuildTag(p, tag, rootDir); err != nil {
			jww.ERROR.Fatalf("could not build properly project: %v", err)
		}
	},
}

func init() {
	addTagFlag(buildCmd)
	addProjectFlag(buildCmd)
	addForceFlag(buildCmd)
	rootCmd.AddCommand(buildCmd)
}

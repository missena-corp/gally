package commands

import (
	"fmt"
	"os/exec"

	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "let Gally run your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		configs := project.UpdatedProjectConfig()
		fmt.Printf("Gally monorepo handler %v", configs)
		for _, c := range configs {
			script, ok := c.Scripts[args[0]]
			if !ok {
				jww.ERROR.Printf("script %s not available", args[0])
				continue
			}
			if _, err := exec.Command(script).Output(); err != nil {
				jww.ERROR.Fatalf("could not run properly script %s: %v", script, err)
			}
		}
	},
}

func init() {
	addRootFolderFlag(runCmd)
	rootCmd.AddCommand(runCmd)
}

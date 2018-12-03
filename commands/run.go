package commands

import (
	"fmt"
	"os/exec"

	"github.com/missena-corp/gally/config"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "let Gally run your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		configs := config.UpdatedProjectConfig()
		for _, c := range configs {
			script, ok := c.Scripts[args[0]]
			if !ok {
				jww.ERROR.Printf("script %s not available", args[0])
			}
			if _, err := exec.Command(script).Output(); err != nil {
				jww.ERROR.Fatalf("could not run properly script %s: %v", script, err)
			}
		}
		fmt.Printf("Gally monorepo handler %v", configs)
	},
}

func init() {
	addRootFolderFlag(runCmd)
	rootCmd.AddCommand(runCmd)
}

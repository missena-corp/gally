package commands

import (
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "let Gally run your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			jww.ERROR.Fatalf("no script provided in command")
		}
		configs := project.UpdatedProjectConfig()
		script := args[0]
		for _, c := range configs {
			out, err := c.Run(script)
			if err != nil {
				jww.ERROR.Fatalf("could not run properly script %s: %v", script, err)
			}
			jww.INFO.Printf(string(out))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

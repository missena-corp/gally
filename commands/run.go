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
		projects := project.FindAllUpdated(rootDir)
		script := args[0]
		for _, p := range projects {
			out, err := p.Run(script)
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

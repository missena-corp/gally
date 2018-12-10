package commands

import (
	"github.com/missena-corp/gally/repo"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build your script for updated files",
	Run: func(cmd *cobra.Command, args []string) {
		handleVerboseFlag()
		tag := repo.Tag()
		if tag == "" {
			jww.INFO.Printf("no tag provided")
			return
		}
		script := args[0]
		for _, p := range projects {
			out, err := p.Run(script)
			if err != nil {
				jww.ERROR.Fatalf("could not build properly script %s: %v", script, err)
			}
			jww.INFO.Printf(string(out))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

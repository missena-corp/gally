package commands

import (
	"fmt"
	"runtime"

	"github.com/missena-corp/gally/helpers"
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Gally version's number",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(gallyVersionString())
			return
		}
		p := project.Find(args[0], rootDir)
		fmt.Println(p.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func gallyVersionString() string {
	program := "Gally, the monorepo slayer"
	version := helpers.Version
	osArch := runtime.GOOS + "/" + runtime.GOARCH
	return fmt.Sprintf("%s %s %s BuildDate: %s", program, version, osArch, helpers.BuildDate)
}

package commands

import (
	"fmt"
	"runtime"

	"github.com/missena-corp/gally/helpers"
	"github.com/missena-corp/gally/project"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Gally version's number",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && projectName == "" {
			fmt.Println(gallyVersionString())
			return
		}
		if projectName == "" {
			jww.WARN.Println("using 'gally version <project>' is deprecated now use 'gally version -p <project>'")
			projectName = args[0]
		}
		p := project.Find(projectName, rootDir)
		fmt.Println(p.Version())
	},
}

func init() {
	addProjectFlag(versionCmd)
	rootCmd.AddCommand(versionCmd)
}

func gallyVersionString() string {
	program := "Gally, the monorepo slayer"
	version := helpers.Version
	osArch := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	return fmt.Sprintf("%s %s %s BuildDate: %s", program, version, osArch, helpers.BuildDate)
}

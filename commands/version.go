package commands

import (
	"fmt"
	"runtime"

	"github.com/missena-corp/gally/helpers"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gally version's number",
	Run: func(cmd *cobra.Command, args []string) {
		printGallyVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printGallyVersion() {
	jww.FEEDBACK.Println(gallyVersionString())
}

func gallyVersionString() string {
	program := "Gally Static Site Generator"
	version := "v" + helpers.Version
	osArch := runtime.GOOS + "/" + runtime.GOARCH
	return fmt.Sprintf("%s %s %s BuildDate: %s", program, version, osArch, helpers.BuildDate)

}

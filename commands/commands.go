package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	rootDir string
	tag     string
	verbose bool
)

func addRootDirFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&rootDir, "root", "r", "", "root directory")
}

func addTagFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "build tag")
}

func addVerboseFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "log level info")
}

func handleVerboseFlag() {
	if verbose {
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelInfo)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

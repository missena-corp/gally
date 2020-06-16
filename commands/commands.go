package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	projectName     string
	rootDir         string
	tag             string
	updatedProjects bool
	verbose         bool
	force           bool
)

func addProjectFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "project name")
}

func addRootDirFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&rootDir, "root", "r", "", "root directory")
}

func addTagFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "build tag")
}

func addUpdatedProjectsFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&projectName, "updated-projects", "u", "", "updated projects")
}

func addVerboseFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "log level info")
}

func addForceFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force the command")
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

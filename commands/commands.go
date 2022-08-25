package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var (
	allProjects   bool
	force         bool
	ignoreMissing bool
	noDependency  bool
	projectName   string
	rootDir       string
	tag           string
	updated       bool
	verbose       bool
)

func addAllProjectsFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&allProjects, "all", "a", false, "all the command")
}

func addIgnoreMissingFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&ignoreMissing, "--ignore-missing", "", false, "ignore missing script")
}

func addForceFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force the command")
}

func addNoDependencyFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&noDependency, "no-dependency", "n", false, "skip dependency")
}

func addProjectFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "project name")
}

func addRootDirFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&rootDir, "root", "r", "", "root directory")
}

func addTagFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "build tag")
}

func addUpdatedFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&updated, "updated", "u", false, "updated projects")
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

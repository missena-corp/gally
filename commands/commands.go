package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
)

func addVerboseFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "log level info")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

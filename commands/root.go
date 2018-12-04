package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var rootCmd = &cobra.Command{
	Use:   "gally",
	Short: "Gally is a monorepo manager",
	Long:  `Lorem ipsum`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check subcommands")
	},
}

func init() {
	addVerboseFlag(rootCmd)
	if verbose {
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelInfo)
	}
}

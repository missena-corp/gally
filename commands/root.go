package commands

import (
	"fmt"

	"github.com/spf13/cobra"
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
	addRootDirFlag(rootCmd)
	addVerboseFlag(rootCmd)
}

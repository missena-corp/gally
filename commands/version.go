package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gally version's number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gally monorepo handler v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

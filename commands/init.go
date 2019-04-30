package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a manifest in the local directory",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := filepath.Abs(".")
		_, name := filepath.Split(path)
		out := fmt.Sprintf("name: %s\n", name) +
			"scripts:\n" +
			fmt.Sprintf("  test: echo \"testing %s 🤞!\"\n", name) +
			"strategies:\n" +
			"  compare-to:\n" +
			"    branch: master\n" +
			"  previous-commit:\n" +
			"    only: master\n" +
			fmt.Sprintf("build: echo \"building %s 💖!\"\n", name) +
			"version: git log -1 --format='%%h' .\n" +
			"tag: false\n"
		if _, err := os.Stat(path + string(filepath.Separator) + ".gally.yml"); os.IsNotExist(err) {
			err := ioutil.WriteFile(path+string(filepath.Separator)+".gally.yml", []byte(out), 0644)
			if err != nil {
				panic(err)
			}
			fmt.Println("File .gally.yml created. Feel free to edit it!")
		} else {
			fmt.Println("File .gally.yml already exists.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

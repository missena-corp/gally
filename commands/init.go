package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a manifest in the local directory",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := filepath.Abs(".")
		_, name := filepath.Split(cwd)
		f := path.Join(cwd, ".gally.yml")
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			fmt.Println("File .gally.yml already exists.")
			return
		}
		config := fmt.Sprintf(`
name: %s
scripts:
	test: echo "no test yet"
build: echo "building %s 💖!"
`, name, name)
		err := ioutil.WriteFile(f, []byte(config), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("File .gally.yml created. Feel free to edit it!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

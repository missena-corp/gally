package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "create a manifest in the local directory",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := filepath.Abs(".")
		if projectName == "" {
			_, projectName = filepath.Split(cwd)
		}
		f := path.Join(cwd, ".gally.yml")
		if _, err := os.Stat(f); !os.IsNotExist(err) && !force {
			fmt.Println("File .gally.yml already exists.")
			return
		}
		config := fmt.Sprintf(`
name: %s
scripts:
	test: echo "no test yet"
build: echo "building %s 💖!"
`, projectName, projectName)
		if err := ioutil.WriteFile(f, []byte(config), 0644); err != nil {
			jww.ERROR.Fatalf("cannot create config file: %v", err)
		}
		fmt.Println("File .gally.yml created. Feel free to edit it!")
	},
}

func init() {
	addForceFlag(initCmd)
	addProjectFlag(initCmd)
	rootCmd.AddCommand(initCmd)
}

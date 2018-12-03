package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const configFileName = ".gally.yml"

type Config struct {
	Ignore     []string
	Scripts    map[string]interface{}
	Strategies map[string]interface{}
}

func FindProjects(rootDir string) (dirs []string, err error) {
	filepath.Walk(rootDir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == configFileName {
			dirs = append(dirs, path.Dir(p))
		}
		return nil
	})
	return
}

func flatSlice(afiles [][]string) []string {
	files := make([]string, 0)
	for _, f := range afiles {
		files = append(files, f...)
	}
	return files
}

func ReadConfig(dir string) (c Config) {
	viper.SetConfigFile(path.Join(dir, configFileName))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return c
}

func UpdatedProject(projects []string, aFiles ...[]string) []string {
	files := flatSlice(aFiles)
	updated := make([]string, 0)
	for _, p := range projects {
		for _, f := range files {
			if strings.HasPrefix(f, p) {
				updated = append(updated, p)
				break
			}
		}
	}
	return updated
}

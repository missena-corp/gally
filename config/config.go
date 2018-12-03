package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
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

func UpdatedProject(projects, files []string) []string {
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

package config

import (
	"os"
	"path"
	"path/filepath"
)

const configFileName = ".gally.yml"

type Config struct {
	Scripts    string
	Strategies string
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

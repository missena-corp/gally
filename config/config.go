package config

import (
	"os"
	"path/filepath"
)

const configFileName = ".gally.yml"

type Config struct {
	Scripts    string
	Strategies string
}

func Find(rootDir string) (confFiles []string, err error) {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == configFileName {
			confFiles = append(confFiles, path)
		}
		return nil
	})
	return
}

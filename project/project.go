package project

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/missena-corp/gally/repo"
	"github.com/spf13/viper"
)

const configFileName = ".gally.yml"

type Config struct {
	Dir        string
	Ignore     []string
	Name       string
	Scripts    map[string]string
	Strategies map[string]Strategy
}

type Strategy struct {
	Branch string
	Only   string
}

func FindProjects(rootDir string) (dirs []string, err error) {
	filepath.Walk(rootDir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == configFileName {
			dirs = append(dirs, path.Dir(p))
		}
		return nil
	})
	return dirs, nil
}

func (c Config) Run(s string) ([]byte, error) {
	script, ok := c.Scripts[s]
	if !ok {
		return nil, fmt.Errorf("script %s not available", s)
	}
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = c.Dir
	return cmd.Output()
}

func ReadConfig(dir string) (c Config) {
	viper.SetConfigFile(path.Join(dir, configFileName))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	c.Dir = dir
	return c
}

func UpdatedFilesByStrategies(strategies map[string]Strategy) []string {
	files := make([]string, 0)
	for name, opts := range strategies {
		switch name {
		case "compare-to":
			res, err := repo.UpdatedFiles(opts.Branch)
			if err != nil {
				continue
			}
			files = append(files, res...)
		case "previous-commit":
			if !repo.IsOnBranch(opts.Only) {
				continue
			}
			res, err := repo.UpdatedFiles("HEAD^1")
			if err != nil {
				continue
			}
			files = append(files, res...)
		default:
			log.Fatalf("unkown strategy %s", name)
		}
	}
	return files
}

func UpdatedProjectConfig() map[string]Config {
	configs := make(map[string]Config)
	projects, _ := FindProjects(repo.Root())
	for _, p := range projects {
		c := ReadConfig(p)
		if WasUpdated(p, c) {
			configs[p] = c
		}
	}
	return configs
}

func WasUpdated(project string, c Config) bool {
	for _, f := range UpdatedFilesByStrategies(c.Strategies) {
		if strings.HasPrefix(f, project) {
			return true
		}
	}
	return false
}

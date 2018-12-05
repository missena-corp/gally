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

type Project struct {
	Dir           string
	ConfigFile    string
	Ignore        []string
	Name          string
	Scripts       map[string]string
	Strategies    map[string]Strategy
	VersionScript string
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

func (p Project) ignored(file string) bool {
	for _, pattern := range p.Ignore {
		files, err := filepath.Glob(pattern)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if f == file {
				return true
			}
		}
	}
	return false
}

func (p Project) include(file string) bool {
	return strings.HasPrefix(file, p.Dir) && !p.ignored(file)
}

func (p Project) Run(s string) ([]byte, error) {
	script, ok := p.Scripts[s]
	if !ok {
		return nil, fmt.Errorf("script %s not available", s)
	}
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = p.Dir
	return cmd.Output()
}

func ReadConfig(dir string) (p Project) {
	v := viper.New()
	v.RegisterAlias("version", "version_script")
	v.SetConfigFile(path.Join(dir, configFileName))
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := v.Unmarshal(&p); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	p.Dir = dir
	return p
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

func UpdatedProjectConfig() map[string]Project {
	configs := make(map[string]Project)
	projects, _ := FindProjects(repo.Root())
	for _, p := range projects {
		c := ReadConfig(p)
		if c.WasUpdated() {
			configs[p] = c
		}
	}
	return configs
}

func (p Project) version() string {
	if p.VersionScript == "" {
		log.Fatalf("no version available in %s", p.Dir)
	}
	cmd := exec.Command("sh", "-c", p.VersionScript)
	cmd.Dir = p.Dir
	v, _ := cmd.Output()
	return string(v)
}

func (p Project) WasUpdated() bool {
	for _, f := range UpdatedFilesByStrategies(p.Strategies) {
		if p.include(f) {
			return true
		}
	}
	return false
}

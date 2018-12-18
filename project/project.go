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
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

const configFileName = ".gally.yml"

type Project struct {
	Build         string
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

func BuildTag(tag string, rootDir string) error {
	sp := strings.Split(tag, "@")
	if len(sp) != 2 {
		return fmt.Errorf("%s is not a valid tag", tag)
	}
	p := Find(sp[0], rootDir)
	if p == nil {
		return fmt.Errorf("project %q not found", sp[0])
	}
	version := p.version()
	if version != "" && version != sp[1] {
		return fmt.Errorf("versions mismatch: %s≠%s", sp[1], version)
	}
	return p.runBuild(version)
}

func (p *Project) exec(str string, env ...string) ([]byte, error) {
	env = append(env, fmt.Sprintf("GALLY_NAME=%s", p.Name))
	cmd := exec.Command("sh", "-c", str)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), env...)
	return cmd.Output()
}

func Find(name string, rootDir string) *Project {
	return FindAll(rootDir)[name]
}

func findPaths(rootDir string) (dirs []string, err error) {
	filepath.Walk(rootDir, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == configFileName {
			dirs = append(dirs, path.Dir(p))
		}
		return nil
	})
	return dirs, nil
}

func FindAll(rootDir string) map[string]*Project {
	if rootDir == "" {
		rootDir = repo.Root()
	}
	projects := make(map[string]*Project)
	paths, _ := findPaths(rootDir)
	for _, path := range paths {
		p := New(path)
		if d, _ := projects[p.Name]; d != nil {
			jww.FATAL.Fatalf("2 projects with name %q exist:\n- %q\n- %q\n", p.Name, d, p.Dir)
		}
		projects[p.Name] = p
	}
	return projects
}

func FindAllUpdated(rootDir string) map[string]*Project {
	projects := make(map[string]*Project)
	for name, project := range FindAll(rootDir) {
		if project.WasUpdated() {
			projects[name] = project
		}
	}
	return projects
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

// Run a command by its name
func (p *Project) Run(s string) error {
	script, ok := p.Scripts[s]
	if !ok {
		return fmt.Errorf("script %s not available", s)
	}
	return p.run(script, fmt.Sprintf("GALLY_VERSION=%s", p.version()))
}

// run a command script for a project
func (p *Project) run(script string, env ...string) error {
	env = append(env, fmt.Sprintf("GALLY_DIR=%s", p.Dir), fmt.Sprintf("GALLY_NAME=%s", p.Name))
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	jww.INFO.Printf("running %q in directory %q\n", script, cmd.Dir)
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func (p Project) runBuild(version string) error {
	return p.run(p.Build, fmt.Sprintf("GALLY_VERSION=%s", version))
}

// New reads current config in directory
// the function is expecting full path as argument
func New(dir string) (p *Project) {
	v := viper.New()
	v.RegisterAlias("version", "version_script")
	v.SetConfigFile(path.Join(dir, configFileName))
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := v.Unmarshal(&p); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	if p.Name == "" {
		p.Name = filepath.Base(dir)
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

func (p Project) version() string {
	if p.VersionScript == "" {
		jww.ERROR.Printf("no version available in %q", p.Dir)
		return ""
	}
	v, _ := p.exec(p.VersionScript)
	return string(v)
}

func (p Project) WasUpdated() bool {
	for _, f := range UpdatedFilesByStrategies(p.Strategies) {
		if strings.HasPrefix(f, p.Dir) && !p.ignored(f) {
			return true
		}
	}
	return false
}

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
	BaseDir       string `mapstructure:"-"`
	BuildScript   string `mapstructure:"build"`
	ConfigFile    string
	Dir           string `mapstructure:"context"`
	Ignore        []string
	Name          string
	Scripts       map[string]string
	Strategies    map[string]Strategy
	VersionScript string `mapstructure:"version"`
}

type Strategy struct {
	Branch string
	Only   string
}

func BuildTag(tag string, rootDir string) error {
	sp := strings.Split(tag, "@")
	if len(sp) != 2 {
		return fmt.Errorf("%q is not a valid tag", tag)
	}
	p := Find(sp[0], rootDir)
	if p == nil {
		return fmt.Errorf("project %q not found", sp[0])
	}
	version := p.Version()
	// this allow to have project without `version` defined
	if version != "" && version != sp[1] {
		return fmt.Errorf("versions mismatch %q≠%q", sp[1], version)
	}
	return p.runBuild(version)
}

func (p *Project) env(env []string) []string {
	return append(
		env,
		fmt.Sprintf("GALLY_CWD=%s", p.Dir),
		fmt.Sprintf("GALLY_NAME=%s", p.Name),
	)
}

func (p *Project) exec(str string, env ...string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", str)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), p.env(env)...)
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
			jww.FATAL.Fatalf("2 projects with name %q exist:\n- %q\n- %q\n", p.Name, d, p.BaseDir)
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

func (p *Project) ignored(file string) bool {
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
		return fmt.Errorf("script %q not available", s)
	}
	return p.run(script, fmt.Sprintf("GALLY_VERSION=%s", p.Version()))
}

// run a command script for a project
func (p *Project) run(script string, env ...string) error {
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), p.env(env)...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	jww.INFO.Printf("running %q in directory %q\n", script, cmd.Dir)
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

func (p *Project) runBuild(version string) error {
	return p.run(
		p.BuildScript,
		fmt.Sprintf("GALLY_TAG=%s@%s", p.Name, version),
		fmt.Sprintf("GALLY_VERSION=%s", version),
	)
}

// New reads current config in directory
// the function is expecting full path as argument
func New(dir string) (p *Project) {
	v := viper.New()
	if !path.IsAbs(dir) {
		d, err := filepath.Abs(dir)
		if err != nil {
			log.Fatalf("unable to expand directory %q: %v", d, err)
		}
		dir = d
	}
	file := path.Join(dir, configFileName)
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file %q: %v", file, err)
	}
	if err := v.Unmarshal(&p); err != nil {
		log.Fatalf("unable to decode file %s into struct: %v", file, err)
	}
	p.BaseDir = dir
	if p.Dir == "" {
		p.Dir = dir
	} else {
		if !path.IsAbs(p.Dir) {
			p.Dir = path.Join(dir, p.Dir)
		}
		if _, err := os.Stat(p.Dir); os.IsNotExist(err) {
			log.Fatalf("context directory %q does not exist", p.Dir)
		}
	}
	if p.Name == "" {
		p.Name = filepath.Base(dir)
	}
	return p
}

func UpdatedFilesByStrategies(strategies map[string]Strategy) []string {
	files := make([]string, 0)
	for name, opts := range strategies {
		switch name {
		case COMPARE_TO:
			res, err := repo.UpdatedFiles(opts.Branch)
			if err != nil {
				continue
			}
			files = append(files, res...)
		case PREVIOUS_COMMIT:
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

func (p *Project) Version() string {
	if p.VersionScript == "" {
		jww.ERROR.Printf("no version available in %q", p.BaseDir)
		return ""
	}
	v, _ := p.exec(p.VersionScript)
	return string(v)
}

func (p *Project) WasUpdated() bool {
	for _, f := range UpdatedFilesByStrategies(p.Strategies) {
		if strings.HasPrefix(f, p.BaseDir) && !p.ignored(f) {
			return true
		}
	}
	return false
}

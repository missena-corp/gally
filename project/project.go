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

var envVars map[string]string

type Project struct {
	BaseDir       string `mapstructure:"-"`
	BuildScript   string `mapstructure:"build"`
	Bumped        *bool
	ConfigFile    string
	Dir           string `mapstructure:"workdir"`
	Disable       bool
	ContextDir    string `mapstructure:"context"`
	Ignore        []string
	Name          string
	RootDir       string
	Scripts       map[string]string
	Strategies    map[string]Strategy
	Tag           *bool `mapstructure:"tag"`
	Updated       *bool
	VersionScript string `mapstructure:"version"`
}

type Projects map[string]*Project

type Strategy struct {
	Branch string
	Only   string
}

func BuildTag(name *string, tag string, rootDir string) error {
	sp := strings.Split(tag, "@")
	if len(sp) != 2 {
		return fmt.Errorf("%q is not a valid tag", tag)
	}
	p := Find(sp[0], rootDir)
	if p == nil {
		return fmt.Errorf("project %q not found", sp[0])
	} else if p != nil && name != nil && p.Name != *name {
		return fmt.Errorf("project %s and tag %q do not match", *name, sp[0])
	}
	version := p.Version()
	// this allow to have project without `version` defined
	if version != "" && version != sp[1] {
		return fmt.Errorf("versions mismatch %q≠%q", sp[1], version)
	}
	return p.runBuild(version)
}

func BuildForceWithTag(name *string, rootDir string) error {
	projects := Projects{}
	if name == nil {
		allProjects := FindAllUpdated(rootDir)
		for _, v := range allProjects {
			if v.wantTag() {
				projects[v.Name] = v
			}
		}
	} else {
		p := Find(*name, rootDir)
		if p == nil {
			return fmt.Errorf("project %q not found", *name)
		}
		if p.wantTag() {
			projects[*name] = p
		}
	}
	for _, p := range projects {
		version := p.Version()
		if version == "" {
			return fmt.Errorf("could not find version for %s", p.Name)
		}
		if err := p.runBuild(p.Version()); err != nil {
			return err
		}
	}
	return nil
}

func BuildNoTag(name *string, rootDir string) error {
	var projects Projects
	if name == nil {
		projects = FindAllUpdated(rootDir)
	} else {
		p := Find(*name, rootDir)
		if p == nil {
			return fmt.Errorf("project %q not found", *name)
		}
		projects[*name] = p
	}
	for _, p := range projects {
		if !p.wantTag() {
			if err := p.runBuild(p.Version()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Project) exec(str string, env Env) ([]byte, error) {
	cmd := exec.Command("sh", "-c", str)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), env.ToSlice()...)
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

func FindAll(rootDir string) Projects {
	if rootDir == "" {
		rootDir = repo.Root()
	}
	projects := make(Projects)
	paths, _ := findPaths(rootDir)
	for _, path := range paths {
		p := New(path, rootDir)
		if d, exists := projects[p.Name]; exists {
			jww.FATAL.Fatalf("2 projects with name %q exist:\n- %q\n- %q\n", p.Name, d.BaseDir, p.BaseDir)
		}
		projects[p.Name] = p
	}
	return projects
}

func FindAllUpdated(rootDir string) Projects {
	projects := make(Projects)
	for name, project := range FindAll(rootDir) {
		if project.WasUpdated() {
			projects[name] = project
		}
	}
	return projects
}

func FindByName(rootDir, name string) *Project {
	projects := FindAll(rootDir)
	return projects[name]
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

func (p *Project) wantTag() bool {
	return p.Tag == nil || *(p.Tag)
}

// New reads current config in directory
// the function is expecting full path as argument
func New(dir, rootDir string) (p *Project) {
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
	if p.Name == "" {
		p.Name = filepath.Base(dir)
	}
	if p.ContextDir != "" && p.Dir == "" {
		fmt.Fprintf(os.Stderr, "**Deprecated** Project %s, replace `context:` with `workdir: in .gally.yml`\n", p.Name)
		p.Dir = p.ContextDir
		p.ContextDir = ""
	}
	p.BaseDir = dir
	if p.Dir == "" {
		p.Dir = dir
	} else {
		if !path.IsAbs(p.Dir) {
			p.Dir = path.Join(dir, p.Dir)
		}
		if _, err := os.Stat(p.Dir); os.IsNotExist(err) {
			log.Fatalf("workdir directory %q does not exist", p.Dir)
		}
	}
	p.RootDir = rootDir
	if !path.IsAbs(rootDir) {
		d, err := filepath.Abs(rootDir)
		if err != nil {
			log.Fatalf("unable to expand directory %q: %v", d, err)
		}
		p.RootDir = d
	}
	return p
}

func newFalse() *bool { b := false; return &b }

func newTrue() *bool { b := true; return &b }

// Run a command by its name
func (p *Project) Run(s string) error {
	script, ok := p.Scripts[s]
	if !ok {
		return fmt.Errorf("script %q not available", s)
	}
	return p.run(script, NewEnv(p))
}

// run a command script for a project
func (p *Project) run(script string, env Env) error {
	if p.Disable {
		jww.WARN.Printf("project %q disabled", p.BaseDir)
		return nil
	}
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = p.Dir
	cmd.Env = append(os.Environ(), env.ToSlice()...)
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
	env := NewEnv(p).Add(Env{
		"GALLY_TAG":     fmt.Sprintf("%s@%s", p.Name, version),
		"GALLY_VERSION": version,
	})
	return p.run(p.BuildScript, env)
}

func (projs Projects) ToSlice() []map[string]interface{} {
	out := make([]map[string]interface{}, 0)
	for _, p := range projs {
		out = append(out, map[string]interface{}{
			"directory":   p.BaseDir,
			"environment": NewCleanEnv(p),
			"name":        p.Name,
			"update":      p.WasUpdated(),
			"version":     p.Version(),
		})
	}
	return out
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
		return repo.Version(p.Dir, p.Ignore)
	}
	v, _ := p.exec(p.VersionScript, NewEnvNoVersion(p))
	return strings.TrimSpace(string(v))
}

func (p *Project) WasBumped() (bumped bool) {
	if p.Bumped != nil {
		return *p.Bumped
	}
	p.Bumped = newFalse()
	if !p.WasUpdated() {
		return false
	}
	currentVersion := p.Version()
	commit := "master"
	if repo.IsOnBranch("master") {
		commit = "HEAD^1"
	}
	repo.Checkout(commit, func() { bumped = p.Version() != currentVersion })
	if bumped {
		p.Bumped = newTrue()
	}
	return bumped
}

func (p *Project) WasUpdated() bool {
	// cache response
	if p.Updated != nil {
		return *p.Updated
	}
	for _, f := range UpdatedFilesByStrategies(p.Strategies) {
		if strings.HasPrefix(f, p.BaseDir) && !p.ignored(f) {
			p.Updated = newTrue()
			return true
		}
	}
	p.Updated = newFalse()
	return false
}

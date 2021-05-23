package project

import (
	"fmt"
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
	// BaseDir is the directory where the configuration file is located
	BaseDir string `mapstructure:"-"`
	// command to build the project
	BuildScript string `mapstructure:"build"`
	// where the config file is located
	ConfigFile string `mapstructure:"-"`
	// list of dependency folders
	DependsOn    []string `mapstructure:"depends_on"`
	Dependencies Dependencies
	// Dir is the directory where the work happen
	Dir     string `mapstructure:"workdir"`
	Disable bool
	Env     []struct {
		Name  string
		Value string
	}
	Ignore        []string
	IsLibrary     bool `mapstructure:"is_library"`
	Name          string
	RootDir       string
	Scripts       map[string]string
	Strategies    Strategies
	Tag           bool   `mapstructure:"tag"`
	VersionScript string `mapstructure:"version"`

	// cache
	bumped  *bool   `mapstructure:"-"`
	updated *bool   `mapstructure:"-"`
	version *string `mapstructure:"-"`
}

type Projects map[string]*Project

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

func BuildForceWithoutTag(name *string, rootDir string, noDep bool) error {
	projects := Projects{}
	if name == nil {
		allProjects := FindAllUpdated(rootDir, noDep)
		for _, p := range allProjects {
			if p.Tag {
				projects[p.Name] = p
			}
		}
	} else {
		p := Find(*name, rootDir)
		if p == nil {
			return fmt.Errorf("project %q not found", *name)
		}
		if p.Tag {
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

func BuildWithoutTag(name *string, rootDir string, noDep bool) error {
	projects := Projects{}
	if name == nil {
		projects = FindAllUpdated(rootDir, noDep)
	} else {
		p := Find(*name, rootDir)
		if p == nil {
			return fmt.Errorf("project %q not found", *name)
		}
		projects[*name] = p
	}
	for _, p := range projects {
		if !p.Tag {
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
	path, _ := os.Getwd()
	abs, _ := filepath.Abs(rootDir)
	base, err := filepath.Rel(path, abs)
	if err != nil {
		base = "."
		err = nil
	}
	cmd := exec.Command("git", "ls-files")
	cmd.Dir = rootDir
	output, _ := cmd.Output()
	for _, v := range strings.Split(string(output), "\n") {
		l := len(v) - len(configFileName)
		if l >= 0 && v[l:] == configFileName {
			if l == 0 {
				dirs = append(dirs, base)
			}
			if l > 0 {
				dirs = append(dirs, fmt.Sprintf("%s%s%s", base, "/", v[:l-1]))
			}
		}
	}
	return
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

func FindAllUpdated(rootDir string, noDep bool) Projects {
	projects := make(Projects)
	for name, project := range FindAll(rootDir) {
		if project.WasUpdated(noDep) {
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

// New reads current config in directory
// the function is expecting full path as argument
func New(dir, rootDir string, isDependency ...bool) (p *Project) {
	v := viper.New()
	if !path.IsAbs(dir) {
		d, err := filepath.Abs(dir)
		if err != nil {
			jww.FATAL.Fatalf("unable to expand directory %q: %v", d, err)
		}
		dir = d
	}
	file := path.Join(dir, configFileName)
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil && len(isDependency) == 0 {
		jww.FATAL.Fatalf("could not read config file %q: %v", file, err)
	}
	if err := v.Unmarshal(&p); err != nil {
		jww.FATAL.Fatalf("unable to decode file %s into struct: %v", file, err)
	}

	// init Name based on current directory if not set in config
	if p.Name == "" {
		p.Name = filepath.Base(dir)
	}

	// init BaseDir
	p.BaseDir = dir
	if p.Dir == "" {
		p.Dir = dir
	}
	if !path.IsAbs(p.Dir) {
		p.Dir = path.Clean(path.Join(dir, p.Dir))
	}
	if _, err := os.Stat(p.Dir); os.IsNotExist(err) {
		jww.FATAL.Fatalf("workdir directory %q does not exist", p.Dir)
	}

	// init RootDir
	p.RootDir = path.Clean(rootDir)
	if !path.IsAbs(rootDir) {
		d, err := filepath.Abs(rootDir)
		if err != nil {
			jww.FATAL.Fatalf("unable to expand directory %q: %v", d, err)
		}
		p.RootDir = d
	}

	// init Strategies, set default strategies if not set
	if len(p.Strategies) == 0 {
		p.Strategies = defaultStrategies
	}

	// init Dependencies
	for _, dep := range p.DependsOn {
		p.Dependencies = append(p.Dependencies, New(path.Join(p.BaseDir, dep), rootDir, true))
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
	return p.run(script, p.env(generateVersion()))
}

// run a command script for a project
func (p *Project) run(script string, env Env) error {
	if p.Disable {
		jww.WARN.Printf("project %q disabled", p.BaseDir)
		return nil
	}
	for _, dep := range p.Dependencies {
		// build dependencies if they are libraries
		if dep.IsLibrary {
			dep.runBuild(dep.Version())
		}
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
	jww.INFO.Printf("building %q version %q\n", p.Name, version)
	return p.run(p.BuildScript, p.env(setVersion(version)))
}

func (projs Projects) ToSlice(noDep bool) []map[string]interface{} {
	out := make([]map[string]interface{}, 0)
	for _, p := range projs {
		addition := map[string]interface{}{
			"directory":   p.BaseDir,
			"environment": p.env(generateVersion()),
			"name":        p.Name,
			"update":      p.WasUpdated(noDep),
			"version":     p.Version(),
		}
		if len(p.DependsOn) != 0 {
			addition["dependencies"] = p.DependsOn
		}
		out = append(out, addition)
	}
	return out
}

func (p *Project) Version() string {
	if p.version != nil {
		return *p.version
	}
	if p.VersionScript == "" {
		version := repo.Version(p.Dir, p.Dependencies.paths(), p.Ignore)
		p.version = &version
		return version
	}
	v, _ := p.exec(p.VersionScript, p.env())
	version := strings.TrimSpace(string(v))
	p.version = &version
	return version
}

func (p *Project) WasUpdated(noDep bool) bool {
	// cache response
	if p.updated != nil {
		return *p.updated
	}
	for _, f := range p.Strategies.UpdatedFiles() {
		if strings.HasPrefix(f, p.BaseDir) && !p.ignored(f) {
			p.updated = newTrue()
			return true
		}
		if noDep {
			continue
		}
		for _, dep := range p.Dependencies {
			if dep.WasUpdated(noDep) {
				p.updated = newTrue()
				return true
			}
		}
	}
	p.updated = newFalse()
	return false
}

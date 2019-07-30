package project

import "fmt"

type Env map[string]string

func (e Env) Add(m map[string]string) Env {
	for k, v := range m {
		e[k] = v
	}
	return e
}

func NewCleanEnv(p *Project) (env Env) {
	env = Env{}
	for _, v := range p.Env {
		env[v.Name] = v.Value
	}
	env["GALLY_PROJECT_WORKDIR"] = p.Dir
	env["GALLY_PROJECT_VERSION"] = p.Version()
	env["GALLY_PROJECT_NAME"] = p.Name
	env["GALLY_PROJECT_ROOT"] = p.BaseDir
	env["GALLY_ROOT"] = p.RootDir
	return
}

func NewEnv(p *Project) (env Env) {
	env = Env{}
	for _, v := range p.Env {
		env[v.Name] = v.Value
	}
	env["GALLY_CWD"] = p.Dir
	env["GALLY_NAME"] = p.Name
	env["GALLY_PROJECT_WORKDIR"] = p.Dir
	env["GALLY_PROJECT_VERSION"] = p.Version()
	env["GALLY_PROJECT_NAME"] = p.Name
	env["GALLY_PROJECT_ROOT"] = p.BaseDir
	env["GALLY_ROOT"] = p.RootDir
	env["GALLY_VERSION"] = p.Version()
	return
}

func NewEnvNoVersion(p *Project) (env Env) {
	env = Env{}
	for _, v := range p.Env {
		env[v.Name] = v.Value
	}
	env["GALLY_CWD"] = p.Dir
	env["GALLY_NAME"] = p.Name
	env["GALLY_PROJECT_WORKDIR"] = p.Dir
	env["GALLY_PROJECT_NAME"] = p.Name
	env["GALLY_PROJECT_ROOT"] = p.BaseDir
	env["GALLY_ROOT"] = p.RootDir
	return
}

func (e Env) ToSlice() []string {
	s := make([]string, 0)
	for k, v := range e {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return s
}

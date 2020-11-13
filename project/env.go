package project

import "fmt"

type Env map[string]string

func (p *Project) env(withVersion bool) (env Env) {
	for _, v := range p.Env {
		env[v.Name] = v.Value
	}
	env["GALLY_PROJECT_WORKDIR"] = p.Dir
	env["GALLY_PROJECT_NAME"] = p.Name
	env["GALLY_PROJECT_ROOT"] = p.BaseDir
	env["GALLY_ROOT"] = p.RootDir
	if withVersion {
		version := p.Version()
		env["GALLY_PROJECT_TAG"] = fmt.Sprintf("%s@%s", p.Name, version)
		env["GALLY_PROJECT_VERSION"] = version
	}
	return env
}

func (e Env) ToSlice() []string {
	s := make([]string, 0)
	for k, v := range e {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return s
}

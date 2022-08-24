package project

import "fmt"

type Env map[string]string
type envOpt func(*Project, *Env)

func generateVersion() envOpt {
	return func(p *Project, e *Env) {
		version := p.Version()
		(*e)["GALLY_PROJECT_VERSION"] = version
		(*e)["GALLY_PROJECT_TAG"] = fmt.Sprintf("%s@%s", p.Name, version)
		(*e)["GALLY_TAG"] = fmt.Sprintf("%s@%s", p.Name, version)
	}
}

func setVersion(version string) envOpt {
	return func(p *Project, e *Env) {
		(*e)["GALLY_PROJECT_VERSION"] = version
		(*e)["GALLY_PROJECT_TAG"] = fmt.Sprintf("%s@%s", p.Name, version)
		(*e)["GALLY_TAG"] = fmt.Sprintf("%s@%s", p.Name, version)
	}
}

func (p *Project) env(opts ...envOpt) Env {
	env := Env{}
	// called first to not overwrite gally based env variables
	for _, v := range p.Env {
		env[v.Name] = v.Value
	}
	env["GALLY_PROJECT_WORKDIR"] = p.WorkDir
	env["GALLY_PROJECT_NAME"] = p.Name
	env["GALLY_PROJECT_ROOT"] = p.BaseDir
	env["GALLY_ROOT"] = p.RootDir
	for _, fn := range opts {
		fn(p, &env)
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

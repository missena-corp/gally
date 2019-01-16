package project

import "fmt"

type Env map[string]string

func (e Env) Add(m map[string]string) Env {
	for k, v := range m {
		e[k] = v
	}
	return e
}

func NewEnv(p *Project) Env {
	return Env{
		"GALLY_CWD":     p.Dir,
		"GALLY_NAME":    p.Name,
		"GALLY_ROOT":    p.RootDir,
		"GALLY_VERSION": p.Version(),
	}
}

func NewEnvNoVersion(p *Project) Env {
	return Env{"GALLY_CWD": p.Dir, "GALLY_ROOT": p.RootDir}
}

func (e Env) ToSlice() []string {
	s := make([]string, 0)
	for k, v := range e {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return s
}
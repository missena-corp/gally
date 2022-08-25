package project

type Dependencies []*Project

func (deps Dependencies) paths() []string {
	paths := make([]string, 0)
	for _, dep := range deps {
		paths = append(paths, dep.WorkDir)
	}
	return paths
}

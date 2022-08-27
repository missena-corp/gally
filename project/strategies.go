package project

import (
	"log"
	"strings"

	"github.com/missena-corp/gally/repo"
)

type Strategy struct {
	Branch string
	Only   string
}

type Strategies map[string]Strategy

const (
	COMPARE_TO      = "compare-to"
	PREVIOUS_COMMIT = "previous-commit"
)

func getDefaultStrategies(branch string) Strategies {
	if branch == "" {
		branch = "master"
	}
	branch = strings.TrimSpace(branch)
	return Strategies{
		COMPARE_TO:      {Branch: branch},
		PREVIOUS_COMMIT: {Only: branch},
	}
}

func (strategies Strategies) UpdatedFiles() []string {
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

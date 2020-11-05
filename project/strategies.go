package project

import (
	"log"

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

var (
	defaultStrategies = Strategies{
		COMPARE_TO: {
			Branch: "master",
		},
		PREVIOUS_COMMIT: {
			Only: "master",
		},
	}
)

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

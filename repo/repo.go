package repo

import (
	"fmt"
	"path"
	"os/exec"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
)

type Repo struct {
	Root string
}

func (r *Repo) UpdatedFiles(branchName string) (files []string, err error) {
	out, err := exec.Command("git", "diff", "--name-only", branchName+"...").Output()
	fmt.Println(string(out))
	if err != nil {
		if _, err = exec.Command("git", "fetch", "origin", branchName).Output(); err != nil {
			jww.ERROR.Printf("cannot get remote branch %s: %v\n", branchName, err)
			return nil, err
		}
		out, err = exec.Command("git", "diff", "--name-only", "FETCH_HEAD...").Output()
		if err != nil {
			jww.ERROR.Printf("cannot make diff for %s: %v\n", branchName, err)
			return nil, err
		}
	}
	for _, f := range strings.Split(string(out), "\n") {
		files = append(files, path.Join(r.Root, strings.TrimSpace(f)))
	}

	return files, nil
}

func NewRepo() (*Repo) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	return &Repo{Root: strings.TrimSpace(string(out))}
}

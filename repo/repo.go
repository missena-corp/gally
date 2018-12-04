package repo

import (
	"os/exec"
	"path"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
)

func BranchExists(branchName string) bool {
	_, err := exec.Command("git", "rev-parse", "--verify", branchName).Output()
	return err == nil
}

func CurrentCommit() string {
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func IsOnBranch(branchName string) bool {
	out, err := exec.Command("git", "branch", "-a", "--contains", CurrentCommit(), "--points-at", branchName).Output()
	return err == nil && strings.TrimSpace(string(out)) != ""
}

func Root() string {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func UpdatedFiles(commit string) (files []string, err error) {
	out, err := exec.Command("git", "diff", "--name-only", commit+"...").Output()
	if err != nil {
		if _, err = exec.Command("git", "fetch", "origin", commit).Output(); err != nil {
			jww.ERROR.Printf("cannot get remote branch %s: %v\n", commit, err)
			return nil, err
		}
		out, err = exec.Command("git", "diff", "--name-only", "FETCH_HEAD...").Output()
		if err != nil {
			jww.ERROR.Printf("cannot make diff for %s: %v\n", commit, err)
			return nil, err
		}
	}
	root := Root()
	for _, f := range strings.Split(string(out), "\n") {
		if f != "" {
			files = append(files, path.Join(root, strings.TrimSpace(f)))
		}
	}

	return files, nil
}

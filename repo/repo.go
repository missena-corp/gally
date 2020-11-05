package repo

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
)

func BranchExists(branchName string) bool {
	_, err := exec.Command("git", "rev-parse", "--verify", branchName).Output()
	return err == nil
}

func Checkout(commit string, fn func()) error {
	if err := exec.Command("git", "checkout", commit).Run(); err != nil {
		if err := exec.Command("git", "fetch", "origin", commit).Run(); err != nil {
			return err
		}
		if err := exec.Command("git", "checkout", "FETCH_HEAD").Run(); err != nil {
			return err
		}
	}
	fn()
	return exec.Command("git", "checkout", "-").Run()
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
		if err := exec.Command("git", "fetch", "origin", commit).Run(); err != nil {
			jww.ERROR.Printf("cannot get remote branch %s: %v\n", commit, err)
			return nil, err
		}
		out, err = exec.Command("git", "diff", "--name-only", "FETCH_HEAD...").Output()
		if err != nil {
			jww.ERROR.Printf("cannot make diff for %s: %v\n", commit, err)
			return nil, err
		}
	}
	if len(out) == 0 && commit != "HEAD^1" {
		return UpdatedFiles("HEAD^1")
	}
	root := Root()
	for _, f := range strings.Split(string(out), "\n") {
		if f != "" {
			files = append(files, path.Clean(path.Join(root, strings.TrimSpace(f))))
		}
	}

	return files, nil
}

func Version(path string, dependencies []string, excluded []string) string {
	args := []string{"log", "-1", "--format=%H", "--", path}
	for _, dep := range dependencies {
		args = append(args, dep)
	}
	for _, ex := range excluded {
		args = append(args, fmt.Sprintf(":(exclude)%s/%s", path, ex))
	}
	cmd := exec.Command("git", args...)
	if out, err := cmd.Output(); err == nil {
		return strings.TrimSpace(string(out))[:16]
	}
	return "unknown"
}

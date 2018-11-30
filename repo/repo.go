package repo

import (
	"fmt"
	"os/exec"
	"strings"

	jww "github.com/spf13/jwalterweatherman"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Repo struct {
	*git.Repository
}

func (r *Repo) UpdatedFiles(from, to *plumbing.Reference) {
	// patch, _ := to.Patch(&from)
	// for i := 0; i < len(patch.Stats()); i++ {
	// 	changedFiles = append(changedFiles, patch.Stats()[i].Name)
	// }

	fromCommit, err := r.CommitObject(from.Hash())
	if err != nil {
		panic(err)
		// return nil, fmt.Errorf("Could not find from commit with hash (%s) %v", from, err)
	}

	toCommit, err := r.CommitObject(to.Hash())
	if err != nil {
		panic(err)
		// return nil, fmt.Errorf("Could not find to commit with hash (%s) %v", to, err)
	}

	diff, err := fromCommit.Patch(toCommit)
	if err != nil {
		panic(err)
		// return nil, fmt.Errorf("Could not generate diff %v", err)
	}
	fmt.Printf("==> %v", diff)

}

func New() (*Repo, error) {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return nil, nil
	}
	return &Repo{r}, nil
}

func (r *Repo) BranchRef(branchName string) *plumbing.Reference {
	if ref, err := r.localBranchRef(branchName); err == nil {
		return ref
	}
	ref, _ := r.remoteBranchRef(branchName)
	return ref
}

func (r *Repo) localBranchRef(branchName string) (*plumbing.Reference, error) {
	head, err := r.Head()
	if err != nil {
		return nil, err
	}
	out, err := exec.Command("git", "merge-base", branchName, head.Hash().String()).Output()
	if err != nil {
		return nil, err
	}
	h, err := r.ResolveRevision(plumbing.Revision(strings.TrimSpace(string(out))))
	if err != nil {
		return nil, err
	}
	return plumbing.NewHashReference(plumbing.NewBranchReferenceName(branchName), *h), nil
}

func (r *Repo) remoteBranchRef(branchName string) (*plumbing.Reference, error) {
	remote, err := r.Remote("origin")
	if err != nil {
		return nil, err
	}
	list, err := remote.List(&git.ListOptions{})
	if err != nil {
		jww.ERROR.Printf("cannot get remote: %v\n", err)
		return nil, err
	}
	for _, ref := range list {
		if ref.Name() == plumbing.NewBranchReferenceName(branchName) {
			return ref, nil
		}
	}
	return nil, fmt.Errorf("%s commit not found", branchName)
}

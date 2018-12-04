package project

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindProjects(t *testing.T) {
	t.Parallel()
	projects, err := FindProjects("..")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expected := []string{"../_examples"}
	if !cmp.Equal(projects, expected) {
		t.Error("projects", projects, "must be equal to", expected)
		t.FailNow()
	}
}

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
		t.Errorf("projects %v must be equal to %v", projects, expected)
		t.FailNow()
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	c := Config{Scripts: map[string]string{"hello": "echo world"}}
	out, err := c.Run("hello")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expected := []byte("world\n")
	if !cmp.Equal(out, expected) {
		t.Errorf("output must be equal to %q but is equal to %q", expected, out)
		t.FailNow()
	}
}

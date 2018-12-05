package project

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindProjectPaths(t *testing.T) {
	t.Parallel()
	paths, err := FindProjectPaths("..")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	expected := []string{"../_examples"}
	if !cmp.Equal(paths, expected) {
		t.Errorf("paths %v must be equal to %v", paths, expected)
		t.FailNow()
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	c := Project{Scripts: map[string]string{"hello": "echo world"}}
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

	expected = []byte("project.go\nproject_test.go\n")
	c = Project{Scripts: map[string]string{"list": "ls"}}
	out, err = c.Run("list")
	if !cmp.Equal(out, expected) {
		t.Errorf("output must be equal to %q but is equal to %q", expected, out)
		t.FailNow()
	}
}

func TestVersion(t *testing.T) {
	t.Parallel()
	c := Project{Dir: "../_examples", VersionScript: "head -1 VERSION"}
	expected := "0.3.5"
	if c.version() != expected {
		t.Errorf("version must be equal to %s", expected)
		t.FailNow()
	}
}

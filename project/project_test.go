package project

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var limitStdout sync.Mutex

func captureOutput(f func()) []byte {
	limitStdout.Lock()
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	limitStdout.Unlock()
	return out
}

func TestBuildVersion(t *testing.T) {
	t.Parallel()
	c := Project{BuildScript: "echo go building $GALLY_VERSION!"}
	out := captureOutput(func() {
		if err := c.runBuild("test"); err != nil {
			t.Error(err)
			t.FailNow()
		}
	})
	expected := []byte("go building test!\n")
	if !cmp.Equal(out, expected) {
		t.Errorf("output must be equal to %q but is equal to %q", expected, out)
		t.FailNow()
	}
}

func TestFindProjectPaths(t *testing.T) {
	t.Parallel()
	paths, err := findPaths("..")
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
	expected := []byte("world\n")
	c := Project{Scripts: map[string]string{"hello": "echo world"}}
	out := captureOutput(func() {
		if err := c.Run("hello"); err != nil {
			t.Error(err)
			t.FailNow()
		}
	})
	if !cmp.Equal(out, expected) {
		t.Errorf("output must be equal to %q but is equal to %q", expected, out)
		t.FailNow()
	}

	expected = []byte("project.go\nproject_test.go\nstrategies.go\n")
	c = Project{Scripts: map[string]string{"list": "ls"}}
	out = captureOutput(func() {
		if err := c.Run("list"); err != nil {
			t.Error(err)
			t.FailNow()
		}
	})
	if !cmp.Equal(out, expected) {
		t.Errorf("output must be equal to %q but is equal to %q", expected, out)
		t.FailNow()
	}
}

func TestVersion(t *testing.T) {
	t.Parallel()
	c := Project{Dir: "../_examples", VersionScript: "head -1 VERSION"}
	expected := "0.3.5"
	if c.Version() != expected {
		t.Errorf("version must be equal to %q but is actually %q", expected, c.Version())
		t.FailNow()
	}
}

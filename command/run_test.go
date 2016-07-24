package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCmdRun(t *testing.T) {
	// Write your code here
}

func TestLoadIgnoredPaths(t *testing.T) {

	dir, err := ioutil.TempDir("", "test-load-ignore-paths")
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer os.RemoveAll(dir)

	if err = os.MkdirAll(filepath.Join(dir, "src"), 0744); err != nil {
		t.Error(err.Error())
		return
	}

	testName := "some_repos"
	target := filepath.Join(dir, "src", GoImportsIgnore)
	err = ioutil.WriteFile(target, []byte(testName), 0644)
	if err != nil {
		t.Error(err.Error())
		return
	}

	res, err := LoadIgnoredPaths(target)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if _, ok := res[testName]; !ok {
		t.Error("LoadIgnoredPaths isn't correct.")
	}

}

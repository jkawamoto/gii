//
// command/run.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

// GoImportsIgnore defines the file name of .goimportsignore.
const GoImportsIgnore = ".goimportsignore"

// CmdRun parses options and run cmdRun.
func CmdRun(c *cli.Context) error {

	if err := cmdRun(c.GlobalString("gopath")); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil

}

func cmdRun(gopath string) (err error) {

	root := filepath.Join(gopath, "src")
	serviceInfos, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}

	ignoreCh := make(chan string)
	pathCh := make(chan string)
	done := make(chan interface{})

	go SearchParallel(pathCh, ignoreCh, done)
	for _, info := range serviceInfos {
		if info.IsDir() {
			pathCh <- filepath.Join(root, info.Name())
		}
	}
	close(pathCh)

	goimportsignore := filepath.Join(root, GoImportsIgnore)
	defined, err := LoadIgnoredPaths(goimportsignore)
	if err != nil {
		return
	}

	fp, err := os.OpenFile(goimportsignore, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	defer fp.Close()

	for {
		select {
		case path := <-ignoreCh:
			rel, err := filepath.Rel(root, path)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if _, exist := defined[rel]; !exist {
				fmt.Fprintln(fp, rel)
				fmt.Printf("Append %s\n", rel)
			}
		case <-done:
			return nil
		}
	}

}

// SearchParallel is a go-routine searching repositories which doesn't belong
// to golang project. It reads search roots from input channel, and write found
// reporitories to output channel. After putting all search roots to input,
// users must close the input channel. When all search finish, the notification
// will be put in done channel.
func SearchParallel(input <-chan string, output chan<- string, done chan<- interface{}) {

	var wg sync.WaitGroup
	for {
		path, ok := <-input
		if !ok {
			break
		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			searchDir(path, output)
		}()
	}
	wg.Wait()
	done <- struct{}{}

}

func searchDir(path string, ch chan<- string) {

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	isRepository := false
	isGoRepository := false
	for _, info := range infos {

		if info.IsDir() {

			if info.Name() == ".git" {
				isRepository = true
			} else if !strings.HasPrefix(".", info.Name()) {

				if isRepository {
					// checking .go folder.
					if searchGoSource(filepath.Join(path, info.Name())) {
						isGoRepository = true
						break
					}
				} else {
					// recursive call
					searchDir(filepath.Join(path, info.Name()), ch)
				}
			}

		} else if strings.HasSuffix(info.Name(), ".go") {
			isGoRepository = true
		}

	}

	if isRepository && !isGoRepository {
		ch <- path
	}

}

func searchGoSource(path string) bool {

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Skip %s (%s)\n", path, err.Error())
		return false
	}

	dirs := map[string]bool{}
	for _, info := range infos {

		if info.IsDir() {
			dirs[info.Name()] = true
		} else {
			if strings.HasSuffix(info.Name(), ".go") {
				return true
			}
		}

	}

	for sub := range dirs {
		if searchGoSource(filepath.Join(path, sub)) {
			return true
		}
	}

	return false

}

// LoadIgnoredPaths reads a .goimportsignore file and returns a set of paths
// included in that file.
func LoadIgnoredPaths(path string) (res map[string]interface{}, err error) {

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	res = map[string]interface{}{}
	for _, path := range strings.Split(string(raw), "\n") {
		res[path] = struct{}{}
	}
	return

}

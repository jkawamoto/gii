# gii
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Build Status](https://travis-ci.org/jkawamoto/gii.svg?branch=master)](https://travis-ci.org/jkawamoto/gii)
[![Code Climate](https://codeclimate.com/github/jkawamoto/gii/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/gii)
[![Release](https://img.shields.io/badge/release-0.1.2-lightgrey.svg)](https://github.com/jkawamoto/gii/releases/tag/v0.1.2)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/e9d135e974a44dade715)

Add not [Go](https://golang.org/) project repositories to `.goimportsignore`.

When you employ go-style directory tree[^1] to maintain all repositories on your
local environment, your `$GOPATH/src` might have many repositories even if they
aren't written in Go, and hence `goimports` becomes slower.

`gii` lists up repositories which aren't Go projects from your `$GOPATH/src`
and writes them to `.goimportsignore` so that `goimports` can ignore those
repositories.

[^1]: http://weblog.bulknews.net/post/89635306479/ghq-pecopercol


## Usage
Run just `gii` if environment variable `$GOPATH` is defined.
If you want to use another root path, use `--gopath` flag.

`gii` appends paths of repositories which aren't Go projects
but haven't been added in `$GOPAH/.goimportsignore` yet.
To delete paths from `.goimportsignore`, edit that file manually.

Here is the help text of `gii`:
~~~
gii [global options]

GLOBAL OPTIONS:
   --gopath GOPATH  GOPATH [$GOPATH]
   --help, -h       show help
   --version, -v    print the version
~~~

## Installation
```sh
$ go get github.com/jkawamoto/gii
```
or if you're a [Homebrew](http://brew.sh/) user,

```sh
$ brew tap jkawamoto/gii
$ brew install gii
```

Compiled binaries are also available in
[Github](https://github.com/jkawamoto/gii/releases).


## License
This software is released under the MIT License, see [LICENSE](LICENSES.md).

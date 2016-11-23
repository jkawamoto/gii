---
---
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/jkawamoto/gii/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/jkawamoto/gii.svg?branch=master)](https://travis-ci.org/jkawamoto/gii)
[![Code Climate](https://codeclimate.com/github/jkawamoto/gii/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/gii)
[![Release](https://img.shields.io/badge/release-0.1.2-lightgrey.svg)](https://github.com/jkawamoto/gii/releases/tag/v0.1.2)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/e9d135e974a44dade715)
<a href="#" data-remodal-target="wallet"><img src="https://img.shields.io/badge/donate-bitcoin-yellow.svg"/></a>

Set repositories which doesn't belong golang project to .goimportsignore.

When you employ go-style directory tree[^1] to maintain all projects
even if such projects are not written in golang,
your `$GOPATH/src` has too many repositories and it makes `goimports` slower.
`gii` lists up repositories which doesn't belong to golang projects from
your `$GOPATH/src` and writes them to `.goimportsignore`,
so that `goimports` doesn't search such repositories.

[^1]: http://weblog.bulknews.net/post/89635306479/ghq-pecopercol


## Usage
Run just `gii` if `$GOPATH` is set. If you want to use another root path,
use `--gopath` flag.

`gii` appends paths of repositories which doesn't belong to golang projects,
and aren't appeared `$GOPAH/.goimportsignore` already.
To delete added paths, please edit `$GOPAH/.goimportsignore` manually.

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
or if you're [Homebrew](http://brew.sh/) user,

```sh
$ brew tap jkawamoto/gii
$ brew install gii
```

Compiled binaries are also available in
Github's [release page](https://github.com/jkawamoto/gii/releases).


# License
This software is released under the MIT License, see [LICENSE](https://github.com/jkawamoto/gii/blob/master/LICENSE).

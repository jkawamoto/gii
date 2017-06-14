---
title: "Gii: Reduce Go import paths"
description: >-
  gii lists up repositories which aren't Go projects from your $GOPATH/src
  and writes them to .goimportsignore so that goimports can ignore those
  repositories.
date: 2016-12-16
lastmod: 2017-06-14
slug: readme
---

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](info/licenses/)
[![Build Status](https://travis-ci.org/jkawamoto/gii.svg?branch=master)](https://travis-ci.org/jkawamoto/gii)
[![Code Climate](https://codeclimate.com/github/jkawamoto/gii/badges/gpa.svg)](https://codeclimate.com/github/jkawamoto/gii)
[![Release](https://img.shields.io/badge/release-0.1.2-lightgrey.svg)](https://github.com/jkawamoto/gii/releases/tag/v0.1.2)
[![Japanese](https://img.shields.io/badge/qiita-%E6%97%A5%E6%9C%AC%E8%AA%9E-brightgreen.svg)](http://qiita.com/jkawamoto/items/e9d135e974a44dade715)
<a href="#" data-remodal-target="wallet"><img src="https://img.shields.io/badge/donate-bitcoin-yellow.svg"/></a>

Add not [Go](https://golang.org/) project repositories to `.goimportsignore`.

When you employ go-style directory tree[^1] to maintain all repositories on your
local environment, your `$GOPATH/src` might have many repositories even if they
aren't written in Go, and hence `goimports` becomes slower.

`gii` lists up repositories which aren't Go projects from your `$GOPATH/src`
and writes them to `.goimportsignore` so that `goimports` can ignore those
repositories.

[^1]: http://weblog.bulknews.net/post/89635306479/ghq-pecopercol


### Usage
Run just `gii` if environment variable `$GOPATH` is defined.
If you want to use another root path, use `--gopath` flag.

`gii` appends paths of repositories which aren't Go projects
but haven't been added in `$GOPAH/.goimportsignore` yet.
To delete paths from `.goimportsignore`, edit that file manually.

Here is the help text of `gii`:
~~~shell
gii [global options]

GLOBAL OPTIONS:
   --gopath GOPATH  GOPATH [$GOPATH]
   --help, -h       show help
   --version, -v    print the version
~~~


### Installation
To build the newest version of Gii, use go get command:

```shell
$ go get github.com/jkawamoto/gii
```

If you're a [Homebrew](http://brew.sh/) user,
you can install Gii by the following commands:

```shell
$ brew tap jkawamoto/gii
$ brew install gii
```

Otherwise, compiled binaries are also available in
[Github](https://github.com/jkawamoto/gii/releases).


### License
This software is released under the MIT License, see [LICENSE](info/licenses/).

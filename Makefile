#
# Makefile
#
# Copyright (c) 2016 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot

default: build

.PHONY: build
build:
	goxc -d=pkg -pv=$(VERSION)

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)

.PHONY: get-deps
get-deps:
	go get -d -t -v .

.PHONY: test
test:
	go test -v ./...

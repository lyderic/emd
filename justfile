alias h := _help

_help:
	@just --list --unsorted --alias-style left --color always \
		--list-heading='' --list-prefix=' ' \
		| sed -e 's/alias: //'

CGO_ENABLED := '0' # otherwise binaries produced are dynamically linked and don't work on musl distros like alpine0

build:
	#!/bin/bash
	version=$(grep VERSION cmd/globals.go | awk -F= '{ print $2 }' | tr -d '"' | sed -e 's/ //')
	echo $version
	GOARCH=amd64 go build -v -o emd-${version}-linux-amd64
	GOARCH=arm64 go build -v -o emd-${version}-linux-arm64
	GOARCH=arm   go build -v -o emd-${version}-linux-arm
	xz --compress --keep -9 --extreme --verbose emd-*-linux-*

tag:
	#!/bin/bash
	git tag 
	tag=$(grep VERSION cmd/globals.go | awk -F= '{ print $2 }' | tr -d '"' | sed -e 's/ /v/')
	echo -n "Please confirm you want to add this tag: [${tag}] [y/N]? "
	read yesno ; [[ "${yesno,,}" =~ ^(y|yes)$ ]] || exit 0
	git tag "${tag}"
	git push origin --tags

[private]
v:
	just --evaluate

set export
set shell := ["bash","-uc"]

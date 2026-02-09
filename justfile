alias h := _help

_help:
	@just --list --unsorted --alias-style left --color always \
		--list-heading='' --list-prefix=' ' \
		| sed -e 's/alias: //'

CGO_ENABLED := '0' # otherwise binaries produced are dynamically linked and don't work on musl distros like alpine0

build:
	GOARCH=amd64 go build -v -o emd_amd64
	GOARCH=arm64 go build -v -o emd_arm64
	GOARCH=arm go build -v -o emd_arm

[private]
v:
	just --evaluate

set export
set shell := ["bash","-uc"]

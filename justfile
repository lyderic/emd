alias h := _help

_help:
	@just --list --unsorted --alias-style left --color always \
		--list-heading='' --list-prefix=' ' \
		| sed -e 's/alias: //'

# smaller binary size
release:
	go build -ldflags="-s -w"
	./emd --version

[private]
v:
	{{justfile()}} --evaluate

set shell := ["bash","-uc"]

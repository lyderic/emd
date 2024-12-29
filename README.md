# emd
Basic markdown viewer for the command line

This viewer is very basic and simple. It uses [Charm's _glamour_ library](https://github.com/charmbracelet/glamour) and does only one thing: render a markdown file and that's it.

## Installation

Download the binary corresponding to your platform. For the moment, I have only tested this on Linux (Arch). If you want paging, `less` needs to be installed.

You can also build it with `go build`.

## Usage

```bash
$ emd file.md
```

## Help

```text
$ emd --help

emd 0.1.6: markdown viewer for the command line

Available themes:
 • dark
 • light
 • dracula
 • pink
 • notty

Usage:
  emd [-n] [-t <theme>] [-w <width>] file.md

Flags:
  -c, --config file   config file (default is $HOME/.emd.yaml)
      --debug         show debugging information
  -h, --help          help for emd
  -n, --no-pager      don't use pager
  -t, --theme name    name of the theme (default "dark")
  -v, --version       version for emd
  -w, --width width   word wrap width

```

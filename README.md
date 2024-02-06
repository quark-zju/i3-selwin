# i3-selwin

Search and jump to a specified window by title. Inspired by [ctrlp.vim](http://kien.github.io/ctrlp.vim/). Require [i3wm](http://i3wm.org/) and [dmenu](http://tools.suckless.org/dmenu/).

## Install

Run:

    go install github.com/quark-zju/i3-selwin@e7ff8c06f62f16f6e96b650c52e69f64bf1a0672

By default the installed binary can be found at `~/go/bin/i3-selwin`.

## Build from source

Clone the repo, then run:

    go build

## Configuration

The program behaves like `dmenu`, accepts all dmenu options. It has no config files. However, you may want to put a line like the following in i3 config:

    bindsym Mod4+m exec i3-selwin -i -b -l 10

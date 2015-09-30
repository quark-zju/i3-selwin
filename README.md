# i3-selwin

Search and jump to a specified window by title. Inspired by [ctrlp.vim](http://kien.github.io/ctrlp.vim/). Require [i3wm](http://i3wm.org/) and [dmenu](http://tools.suckless.org/dmenu/).

## Build

Use standard go tool with [vendor support](https://golang.org/s/go15vendor):

    export GO15VENDOREXPERIMENT=1
    go get gitcafe.com/quark/i3-selwin

## Configuration

The program behaves like `dmenu`, accepts all dmenu options. It has no config files. However, you may want to put a line like the following in i3 config:

    bindsym Mod4+m exec i3-selwin -i -b -l 10

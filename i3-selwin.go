package main

import (
	"fmt"
	"github.com/proxypoke/i3ipc"
	"os"
	"os/exec"
	"strings"
)

func dfsTree(t *i3ipc.I3Node) (names []string, nodes []*i3ipc.I3Node) {
	if t.Geometry.Width > 0 && t.Geometry.Height > 0 && t.Window > 0 && !t.Focused {
		names = append(names, t.Name)
		nodes = append(nodes, t)
	}
	visitChildren := func(children *[]i3ipc.I3Node) {
		for i, _ := range *children {
			newNames, newNodes := dfsTree(&(*children)[i])
			names = append(names, newNames...)
			nodes = append(nodes, newNodes...)
		}
	}
	visitChildren(&t.FloatingNodes)
	visitChildren(&t.Nodes)
	return
}

func checkError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(0)
	}
}

func main() {
	ipc, e := i3ipc.GetIPCSocket()
	checkError(e)

	tree, e := ipc.GetTree()
	checkError(e)

	names, nodes := dfsTree(&tree)

	cmd := exec.Command("dmenu", os.Args[1:len(os.Args)]...)
	cmd.Stdin = strings.NewReader(strings.Join(names, "\n"))

	out, e := cmd.Output()
	checkError(e)

	name := strings.TrimRight(string(out), "\n")

	var node *i3ipc.I3Node
	for i, s := range names {
		if s == name {
			node = nodes[i]
			break
		}
	}

	if node != nil {
		msg := fmt.Sprint("[con_id=", node.Id, "] focus")
		ipc.Raw(i3ipc.I3Command, msg)
	}
}

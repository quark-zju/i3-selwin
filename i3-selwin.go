package main

import (
	"fmt"
	"github.com/quark-zju/i3-selwin/i3ipc"
	"os"
	"os/exec"
	"strings"
)

func dfsTree(t *i3ipc.I3Node) (names []string, nodes []*i3ipc.I3Node) {
	if t.Layout == "dockarea" {
		return
	}

	if t.Geometry.Width > 0 && t.Geometry.Height > 0 && t.Window > 0 && !t.Focused {
		name := t.Name
		if len(name) == 0 {
			name = fmt.Sprint("(Unnamed) ", t.Window)
		}
		names = append(names, name)
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
		os.Exit(1)
	}
}

func main() {
	ipc, e := i3ipc.GetIPCSocket()
	checkError(e)

	tree, e := ipc.GetTree()
	checkError(e)

	var node *i3ipc.I3Node
	names, nodes := dfsTree(&tree)
	switch len(names) {
	case 0:
		os.Exit(0)
	case 1:
		node = nodes[0]
	default:
		cmd := exec.Command("dmenu", os.Args[1:len(os.Args)]...)
		cmd.Stdin = strings.NewReader(strings.Join(names, "\n"))

		out, e := cmd.Output()
		checkError(e)

		name := strings.TrimRight(string(out), "\n")

		for i, s := range names {
			if s == name {
				node = nodes[i]
				break
			}
		}
	}

	if node != nil {
		msg := fmt.Sprint("[con_id=\"", node.Id, "\"] focus")
		fmt.Print("Sent: ", msg, "\n")
		reply, err := ipc.Raw(i3ipc.I3Command, msg)
		checkError(err)
		s := string(reply)
		fmt.Print("Received: ", s, "\n")
	}
}

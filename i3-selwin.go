package main

import (
	"fmt"
	"github.com/proxypoke/i3ipc"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Candidate struct {
	Name string
	Id   int32
}

func dfsTree(t *i3ipc.I3Node) (cands []*Candidate) {
	if t.Geometry.Width > 0 && t.Geometry.Height > 0 && t.Window > 0 {
		cands = append(cands, &Candidate{Name: t.Name, Id: t.Id})
	}
	for _, c := range t.Nodes {
		cands = append(cands, dfsTree(&c)...)
	}
	return
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	ipc, e := i3ipc.GetIPCSocket()
	checkError(e)

	n, e := ipc.GetTree()
	checkError(e)

	cands := dfsTree(&n)

	var input string
	for _, c := range cands {
		input += c.Name + "\n"
	}

	cmd := exec.Command("dmenu", os.Args[1:len(os.Args)]...)
	cmd.Stdin = strings.NewReader(input)

	out, e := cmd.Output()
	checkError(e)

	name := strings.TrimRight(string(out), "\n")

	var cand *Candidate
	for _, c := range cands {
		if c.Name == name {
			cand = c
			break
		}
	}

	if cand != nil {
		msg := fmt.Sprint("[con_id=", cand.Id, "] focus")
		ipc.Raw(i3ipc.I3Command, msg)
	}
}

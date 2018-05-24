package main

import (
	"fmt"

	"github.com/tmadeira/spnbn/add"
	"github.com/tmadeira/spnbn/spn"
	"github.com/tmadeira/spnbn/zhao"
)

func recPrint(u add.Node, level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}

	fmt.Printf("%s\n", u.Label())
	for _, v := range u.Children() {
		recPrint(v, level+1)
	}
}

func makeSPN(x1, x2 int) spn.Node {
	S := spn.NewSumNode()
	a, b, c := spn.NewProductNode(), spn.NewProductNode(), spn.NewProductNode()
	d, e := spn.NewTerminalNode(x1, 0.6), spn.TerminalNode(x1, 0.9)
	f, g := spn.NewTerminalNode(x2, 0.3), spn.NewTerminalNode(x2, 0.2)

	a.AddChild(d)
	a.AddChild(f)
	b.AddChild(d)
	b.AddChild(g)
	c.AddChild(e)
	c.AddChild(g)
	S.AddChild(a, 20.0/35.0)
	S.AddChild(b, 6.0/35.0)
	S.AddChild(c, 9.0/35.0)

	return S
}

func main() {
	S := makeSPN(1, 2)
	B := zhao.BuildBN(S)

	fmt.Printf("Hidden variables = %d\n", len(B.Hid))
	fmt.Printf("Observable variables = %d\n\n", len(B.Obs))

	fmt.Printf("SPN -> BN observable variables mapping:\n")
	for k, v := range B.Obs {
		fmt.Printf("* X_%d -> %d\n", k, v)
	}
	fmt.Printf("\n")

	fmt.Printf("BN structure (adjacency lists):\n")
	for _, u := range B.Hid {
		fmt.Printf("* %d -> %v\n", u, B.Edges[u])
	}
	fmt.Printf("\n")

	fmt.Printf("ADDs:\n")
	for i := 0; i < len(B.ADDs); i++ {
		fmt.Printf("* %d:\n", i)
		recPrint(B.ADDs[i], 1)
	}
}

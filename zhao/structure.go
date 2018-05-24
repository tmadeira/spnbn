package zhao

import (
	"github.com/tmadeira/spnbn/bn"
	"github.com/tmadeira/spnbn/spn"
)

func BuildStructure(S spn.Node, B *bn.BN, built map[spn.Node]bool) {
	if built[S] {
		return
	}

	built[S] = true

	if S.Type() == spn.Terminal {
		v := S.(*spn.TerminalNode)
		B.AddObservableVariable(v.Label)
		return
	}

	for _, ch := range S.Children() {
		BuildStructure(ch, B, built)
	}

	if S.Type() == spn.Sum {
		node := B.AddHiddenVariable(S)
		for _, v := range S.Scope() {
			B.AddEdge(node, v)
		}
	}
}

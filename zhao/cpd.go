package zhao

import (
	"github.com/tmadeira/spnbn/add"
	"github.com/tmadeira/spnbn/bn"
	"github.com/tmadeira/spnbn/spn"
)

func BuildADD(S spn.Node, X int, B *bn.BN, built map[spn.Node]add.Node) add.Node {
	if add, ok := built[S]; ok {
		return add
	}

	if S.Type() == spn.Terminal {
		built[S] = add.MakeDecisionStump(B.Obs[X], S)
		return built[S]
	}

	if S.Type() == spn.Sum {
		a := add.NewInternalNode(B.Hid[S])
		for _, ch := range S.Children() {
			a.AddChild(BuildADD(ch, X, B, built))
		}
		built[S] = a
		return built[S]
	}

	for _, ch := range S.Children() {
		for _, v := range ch.Scope() {
			if v == X {
				built[S] = BuildADD(ch, X, B, built)
				return built[S]
			}
		}
	}

	return nil
}

func BuildObservableVariablesADDs(S spn.Node, B *bn.BN) {
	for _, x := range S.Scope() {
		built := make(map[spn.Node]add.Node)
		B.SetADD(x, BuildADD(S, x, B, built))
	}
}

package zhao

import (
	"github.com/tmadeira/spnbn/bn"
	"github.com/tmadeira/spnbn/spn"
)

func BuildBN(S spn.Node) *bn.BN {
	B := bn.NewBN()

	built := make(map[spn.Node]bool)
	BuildStructure(S, B, built)

	BuildObservableVariablesADDs(S, B)

	return B
}

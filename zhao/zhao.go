// Package zhao implements Zhao et al. (2015) algorithm to convert SPNs into BNs with ADDs.
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

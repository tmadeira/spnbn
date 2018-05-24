// Package bn implements Bayesian Networks with CPDs represented by algebraic decision diagrams.
package bn

import (
	"github.com/tmadeira/spnbn/add"
	"github.com/tmadeira/spnbn/spn"
)

type BN struct {
	Obs   map[int]int
	Hid   map[spn.Node]int
	N     int
	Edges [][]int
	ADDs  []add.Node
}

func NewBN() *BN {
	b := &BN{}
	b.Obs = make(map[int]int)
	b.Hid = make(map[spn.Node]int)
	b.Edges = [][]int{}
	b.ADDs = []add.Node{}
	return b
}

func (b *BN) newNode() int {
	v := b.N
	b.Edges = append(b.Edges, []int{})
	b.ADDs = append(b.ADDs, nil)
	b.N++
	return v
}

func (b *BN) newEdge(u, v int) {
	b.Edges[u] = append(b.Edges[u], v)
}

func (b *BN) AddObservableVariable(label int) int {
	v, ok := b.Obs[label]
	if !ok {
		v = b.newNode()
		b.Obs[label] = v
	}

	return v
}

func (b *BN) AddHiddenVariable(s spn.Node) int {
	v := b.newNode()
	b.Hid[s] = v
	b.ADDs[v] = add.MakeDecisionStump(v, s)
	return v
}

func (b *BN) AddEdge(u, vLabel int) {
	v := b.Obs[vLabel]
	b.newEdge(u, v)
}

func (b *BN) SetADD(label int, a add.Node) {
	v := b.Obs[label]
	b.ADDs[v] = a
}

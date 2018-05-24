package add

import (
	"strconv"

	"github.com/tmadeira/spnbn/spn"
)

type Node interface {
	Children() []Node
	Label() string
}

type InternalNode struct {
	variable int
	ch       []Node
}

func (v *InternalNode) Children() []Node {
	return v.ch
}

func (v *InternalNode) AddChild(w Node) {
	v.ch = append(v.ch, w)
}

func (v *InternalNode) Label() string {
	return "Y_" + strconv.FormatInt(int64(v.variable), 10)
}

func NewInternalNode(variable int) *InternalNode {
	return &InternalNode{variable, []Node{}}
}

type TerminalNode struct {
	value float64
}

func (v *TerminalNode) Children() []Node {
	return []Node{}
}

func (v *TerminalNode) Label() string {
	return strconv.FormatFloat(v.value, 'f', 2, 64)
}

func MakeDecisionStump(label int, S spn.Node) Node {
	if S.Type() == spn.Sum {
		v := S.(*spn.SumNode)
		add := &InternalNode{label, []Node{}}
		for _, w := range v.W {
			add.ch = append(add.ch, &TerminalNode{w})
		}
		return add
	}

	v := S.(*spn.TerminalNode)
	add := &InternalNode{label, []Node{}}
	add.ch = append(add.ch, &TerminalNode{v.Value})
	add.ch = append(add.ch, &TerminalNode{1 - v.Value})
	return add
}

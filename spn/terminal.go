package spn

type TerminalNode struct {
	DefaultNode
	Label int
	Value float64
}

func (s *TerminalNode) Type() NodeType {
	return Terminal
}

func NewTerminalNode(label int, value float64) *TerminalNode {
	sc := make(map[int]bool)
	sc[label] = true
	return &TerminalNode{DefaultNode{[]Node{}, sc}, label, value}
}

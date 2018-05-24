package spn

type SumNode struct {
	DefaultNode
	W []float64
}

func (s *SumNode) Type() NodeType {
	return Sum
}

func (s *SumNode) AddChild(v Node, w float64) {
	s.ch = append(s.ch, v)
	s.W = append(s.W, w)

	for _, w := range v.Scope() {
		s.sc[w] = true
	}
}

func NewSumNode() *SumNode {
	return &SumNode{DefaultNode{[]Node{}, make(map[int]bool)}, []float64{}}
}

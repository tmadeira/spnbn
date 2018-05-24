package spn

type ProductNode struct {
	DefaultNode
}

func (s *ProductNode) Type() NodeType {
	return Product
}

func (s *ProductNode) AddChild(v Node) {
	s.ch = append(s.ch, v)
	for _, w := range v.Scope() {
		s.sc[w] = true
	}
}

func NewProductNode() *ProductNode {
	return &ProductNode{DefaultNode{[]Node{}, make(map[int]bool)}}
}

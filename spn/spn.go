package spn

type NodeType int

const (
	Terminal NodeType = 0
	Sum      NodeType = 1
	Product  NodeType = 2
)

type Node interface {
	Type() NodeType
	Children() []Node
	Scope() []int
}

type DefaultNode struct {
	ch []Node
	sc map[int]bool
}

func (s *DefaultNode) Children() []Node {
	return s.ch
}

func (s *DefaultNode) Scope() []int {
	keys := make([]int, len(s.sc))

	i := 0
	for k := range s.sc {
		keys[i] = k
		i++
	}

	return keys
}

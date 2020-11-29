package comparejson

type List interface {
	Size() int
	Contains(element *Node) bool
	ToArray() []*Node
}

type NodeSet struct {
	hashSet map[*Node]bool
}

func (n *NodeSet) Size() int {
	return len(n.hashSet)
}

func (n *NodeSet) Contains(element *Node) bool {
	for k, _ := range n.hashSet {
		if DeepEqualNodes(k, element) {
			return true
		}
	}
	return false
}

func (n *NodeSet) ToArray() []*Node {
	result := []*Node{}
	for k, _ := range n.hashSet {
		result = append(result, k)
	}
	return result
}

func NewNodeSet() List {
	return &NodeSet{}
}

type Node struct {
	Name       string
	IsList     bool
	Attributes map[string]string
	Children   map[string]*Node
	List       List
	Parent     *Node
}

const MisleadingErrorMessage = "Don`t call deep-equal on the same element! It is misleading!"

func DeepEqualNodes(n1, n2 *Node) bool {
	if n1 == n2 {
		panic(MisleadingErrorMessage)
	}

	if n1.Name != n2.Name {
		return false
	}

	if n1.IsList != n2.IsList {
		return false
	}

	if len(n1.Attributes) != len(n2.Attributes) {
		return false
	}

	for k, v1 := range n1.Attributes {
		v2 := n2.Attributes[k]

		if v1 != v2 {
			return false
		}
	}

	for k, v2 := range n2.Attributes {
		v1 := n1.Attributes[k]

		if v1 != v2 {
			return false
		}
	}

	if len(n1.Children) != len(n2.Children) {
		return false
	}

	for k, c2 := range n2.Children {
		c1 := n1.Children[k]

		if DeepEqualNodes(c1, c2) {
			return false
		}
	}

	if !DeepEqualLists(n1.List, n2.List) {
		return false
	}

	return true

}

func DeepEqualLists(l1 List, l2 List) bool {
	if l1 == l2 {
		panic(MisleadingErrorMessage)
	}

	if l1.Size() != l2.Size() {
		return false
	}

	l1Array := l1.ToArray()

	for _, e1 := range l1Array {
		if !l2.Contains(e1) {
			return false
		}
	}

	return true

}

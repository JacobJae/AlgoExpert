package graphs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array2 := []string{n.Name}
	for _, child := range n.Children {
		array2 = append(array2, child.DepthFirstSearch(array2)...)
	}
	return array2
}

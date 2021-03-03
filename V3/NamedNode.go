package V3

import "fmt"

type INamedNode interface {
	INode
	getName() string
}

type NamedNode struct {
	*Node
	name string
}

func NewNamedNode(node *Node, name string) *NamedNode {
	node.definer = node.definer + name
	return &NamedNode{
		Node: node,
		name: name,
	}
}

func (n *NamedNode) getName() string {
	return n.name
}

func (n *NamedNode) copy() INode {
	copy := NewNamedNode(NewNode(n.typeId, n.rank, n.maxChilds), n.name)
	copy.childs = make([]INode, len(n.childs))

	for i, child := range n.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (n *NamedNode) print() {
	fmt.Print(n.name)
	n.Node.print()
}
func (n *NamedNode) printTree(indetation int) {
	printIndentation(indetation)
	fmt.Print(n.name)
	n.Node.printTree(indetation)
}

var solvableTermNodes []INamedNode

func setUpNamedNodeSlice() {
	for _, x := range mathOperators {
		solvableTermNodes = append(solvableTermNodes, x)
	}

	for _, x := range mathFunctions {
		solvableTermNodes = append(solvableTermNodes, x)
	}

}

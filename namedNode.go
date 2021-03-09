package AlgebraCalculator

import "AlgebraCalculator/log"

type iNamedNode interface {
	iNode
	getName() string
}

type namedNode struct {
	*node
	name string
}

func newNamedNode(node *node, name string) *namedNode {
	node.definer = node.definer + name
	return &namedNode{
		node: node,
		name: name,
	}
}

func (n *namedNode) getName() string {
	return n.name
}

func (n *namedNode) copy() iNode {
	copy := newNamedNode(newNode(n.typeId, n.rank, n.maxChilds), n.name)
	copy.childs = make([]iNode, len(n.childs))

	for i, child := range n.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (n *namedNode) print() {
	log.Print(n.name)
	n.node.print()
}
func (n *namedNode) printTree(indetation int) {
	printIndentation(indetation)
	log.Print(n.name)
	n.node.printTree(indetation)
}

var solvableTermNodes []iNamedNode

func initNamedNodeSlice() {
	for _, x := range mathOperators {
		solvableTermNodes = append(solvableTermNodes, x)
	}

	for _, x := range mathFunctions {
		solvableTermNodes = append(solvableTermNodes, x)
	}

	for _, x := range termFunctions {
		solvableTermNodes = append(solvableTermNodes, x)
	}

	solvableTermNodes = append(solvableTermNodes, newSubOpperaton())

}

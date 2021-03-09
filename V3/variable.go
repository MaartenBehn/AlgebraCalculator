package V3

type variable struct {
	*namedNode
}

func newVariable(name string) *variable {
	return &variable{
		namedNode: newNamedNode(newNode(typVariable, rankNotSolvable, 0), name),
	}
}

func (v *variable) copy() iNode {
	copy := newVariable(v.name)
	copy.childs = make([]iNode, len(v.childs))

	for i, child := range v.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}

package V3

type Variable struct {
	*NamedNode
}

func NewVariable(name string) *Variable {
	return &Variable{
		NamedNode: NewNamedNode(NewNode(TypVariable, RankNotSolvable, 0), name),
	}
}

func (v *Variable) copy() INode {
	copy := NewVariable(v.name)
	copy.childs = make([]INode, len(v.childs))

	for i, child := range v.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}

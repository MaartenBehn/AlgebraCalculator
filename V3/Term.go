package V3

import "fmt"

type Term struct {
	*NamedNode
	variables []*Variable
}

func NewTerm(name string, variables []*Variable) *Term {
	return &Term{
		NamedNode: NewNamedNode(NewNode(TypTerm, RankTerm, 1), name),
		variables: variables,
	}
}

func (t *Term) copy() INode {
	copy := NewTerm(t.name, t.variables)
	copy.childs = make([]INode, len(t.childs))

	for i, child := range t.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (t *Term) print() {
	fmt.Print(t.name)
	if len(t.variables) > 0 {

		fmt.Print("<")
		for i, variable := range t.variables {
			variable.print()
			if i < len(t.variables)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print(">")
	}
	fmt.Print(" = ")

	t.Node.print()
}

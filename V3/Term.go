package V3

import "fmt"

type Term struct {
	*NamedNode
	variables []*Variable
	root      INode
}

func NewTerm(name string, variables []*Variable) *Term {
	return &Term{
		NamedNode: NewNamedNode(NewNode(TypTerm, RankTerm, 1), name),
		variables: variables,
	}
}

func (t *Term) solve() bool {
	t.Node.solve()

	t.replaceTermVariables(t.root, t.childs)
	firstNode := t.root.getChilds()[0]
	insertNode(t, firstNode)
	return true
}
func (t *Term) replaceTermVariables(node INode, replacements []INode) {
	for _, child := range node.getChilds() {
		t.replaceTermVariables(child, replacements)
	}

	for i, variable := range t.variables {
		if node.getDefiner(false) == variable.getDefiner(false) {
			insertNode(node, replacements[i].copy())
		}
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
	copy.root = t.root.copy()
	return copy
}
func (t *Term) print() {
	fmt.Print(t.name)
	t.Node.print()
}

func (t *Term) printTerm() {
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

	t.root.print()
}

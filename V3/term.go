package V3

import (
	"AlgebraCalculator/log"
	"strings"
)

type term struct {
	*namedNode
	variables []*variable
	root      iNode
	logParts  []string
}

func newTerm(name string, variables []*variable) *term {
	return &term{
		namedNode: newNamedNode(newNode(typTerm, rankTerm, 1), name),
		variables: variables,
	}
}

func (t *term) solve() bool {
	t.node.solve()

	t.replaceTermVariables(t.root, t.childs)
	firstNode := t.root.getChilds()[0]
	insertNode(t, firstNode)
	return true
}
func (t *term) replaceTermVariables(node iNode, replacements []iNode) {
	for _, child := range node.getChilds() {
		t.replaceTermVariables(child, replacements)
	}

	for i, variable := range t.variables {
		if node.getDefiner(false) == variable.getDefiner(false) {
			insertNode(node, replacements[i].copy())
		}
	}
}

func (t *term) copy() iNode {
	copy := newTerm(t.name, t.variables)
	copy.childs = make([]iNode, len(t.childs))

	for i, child := range t.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	copy.root = t.root.copy()
	return copy
}
func (t *term) print() {
	log.Print(t.name)
	t.node.print()
}

func (t *term) printTerm() {
	log.Print(t.name)
	if len(t.variables) > 0 {

		log.Print("<")
		for i, variable := range t.variables {
			variable.print()
			if i < len(t.variables)-1 {
				log.Print(" ")
			}
		}
		log.Print(">")
	}
	log.Print(" = ")

	t.root.print()
}

func (t *term) addLog(texts []string) {
	t.logParts = append(t.logParts, texts...)
}
func (t *term) clearLog() {
	t.logParts = nil
}
func (t *term) getLogString() string {
	builder := strings.Builder{}
	for _, string := range t.logParts {
		builder.WriteString(string)
	}
	return builder.String()
}

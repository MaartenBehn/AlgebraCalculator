package AlgebraCalculator

import (
	"AlgebraCalculator/log"
)

type term struct {
	*namedNode
	variables []*variable
	root      iNode
	fromLine  int
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
func (t *term) solveTerm() {

	log.Print("\n")
	t.printTerm()
	log.Print("\n")

	root := t.root
	log.Print("Parsed:")
	root.printTree(0)
	log.Print("\n")

	for _, ruleList := range simpRules {
		if ruleList == nil {
			continue
		}

		run := true
		for i := 0; i < 1000 && run; i++ {

			run = false
			run2 := true
			for run2 {
				run2 = root.sort()

				if run2 {
					log.Print("Sort:")
					root.printTree(0)
					log.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				run2 = root.solve()

				if run2 {
					log.Print("Solve:")
					root.printTree(0)
					log.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				for _, rule := range ruleList {
					run2 = rule.applyRule(root, rule.search)

					if run2 {
						log.Print("Simpify: ")
						log.Printf("%s", rule.base)
						root.printTree(0)
						log.Print("\n")

						root.solve()

						run = true
						break
					}
				}
			}
		}
	}
	t.root = root

	log.Print(" => ")
	t.printTerm()
	log.Print("\n")
}

package V3

import (
	"AlgebraCalculator/log"
)

type subOpperation struct {
	*namedNode
}

func newSubOpperaton() *subOpperation {
	return &subOpperation{
		namedNode: newNamedNode(newNode(typSubOperation, rankSubOpperation, 2), "."),
	}
}

func (s *subOpperation) copy() iNode {
	copy := newSubOpperaton()
	copy.childs = make([]iNode, len(s.childs))

	for i, child := range s.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (s *subOpperation) solve() bool {
	s.node.solve()

	if false {
		if len(s.childs) == 2 && s.childs[0].getType() == typVector && s.childs[1].getType() == typVector {
			result := newVector(nil)

			number := int64(s.childs[1].(*vector).values[0])
			for number > 0 {
				digit := number % 10
				number /= 10

				result.append(newVector([]float64{s.childs[0].(*vector).values[digit]}))
			}
			replaceNode(s.childs[0], result)
			insertNode(s, s.childs[0])
		}
	} else {
		s.solveReplace(s.childs[0])
		insertNode(s, s.childs[0])
	}
	return true
}
func (s *subOpperation) solveReplace(node iNode) {
	for _, child := range node.getChilds() {
		s.solveReplace(child)
	}

	if node.getType() == typVector {
		result := newVector(nil)
		number := int64(s.childs[1].(*vector).values[0])
		for number > 0 {
			digit := number % 10
			number /= 10

			result.append(newVector([]float64{node.(*vector).values[digit]}))
		}
		replaceNode(node, result)
	} else if node.getType() == typVariable {

		subOperation := newSubOpperaton()
		configuration := s.childs[1].copy()

		insertNode(node, subOperation)
		subOperation.setChilds([]iNode{node, configuration})
		node.setParent(subOperation)
		configuration.setParent(subOperation)
	}
}

func (s *subOpperation) print() {
	log.Print("(")
	if len(s.childs) < 1 {
		return
	}
	s.childs[0].print()
	log.Printf(" %s ", s.name)
	if len(s.childs) < 2 {
		return
	}
	s.childs[1].print()
	log.Print(")")
}
func (s *subOpperation) printTree(indentation int) {
	if len(s.childs) < 1 {
		return
	}
	s.childs[0].printTree(indentation + 1)

	printIndentation(indentation)
	log.Printf("%s\n", s.name)
	if len(s.childs) < 2 {
		return
	}
	s.childs[1].printTree(indentation + 1)
}

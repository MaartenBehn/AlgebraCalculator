package V3

import (
	"fmt"
)

type SubOpperation struct {
	*NamedNode
}

func NewSubOpperaton() *SubOpperation {
	return &SubOpperation{
		NamedNode: NewNamedNode(NewNode(TypSubOperation, RankSubOpperation, 2), "."),
	}
}

func (s *SubOpperation) copy() INode {
	copy := NewSubOpperaton()
	copy.childs = make([]INode, len(s.childs))

	for i, child := range s.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (s *SubOpperation) solve() bool {
	s.Node.solve()

	if false {
		if len(s.childs) == 2 && s.childs[0].getType() == TypVector && s.childs[1].getType() == TypVector {
			result := NewVector(nil)

			number := int64(s.childs[1].(*Vector).values[0])
			for number > 0 {
				digit := number % 10
				number /= 10

				result.append(NewVector([]float64{s.childs[0].(*Vector).values[digit]}))
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
func (s *SubOpperation) solveReplace(node INode) {
	for _, child := range node.getChilds() {
		s.solveReplace(child)
	}

	if node.getType() == TypVector {
		result := NewVector(nil)
		number := int64(s.childs[1].(*Vector).values[0])
		for number > 0 {
			digit := number % 10
			number /= 10

			result.append(NewVector([]float64{node.(*Vector).values[digit]}))
		}
		replaceNode(node, result)
	} else if node.getType() == TypVariable {

		subOperation := NewSubOpperaton()
		configuration := s.childs[1].copy()

		insertNode(node, subOperation)
		subOperation.setChilds([]INode{node, configuration})
		node.setParent(subOperation)
		configuration.setParent(subOperation)
	}
}

func (s *SubOpperation) print() {
	fmt.Print("(")
	if len(s.childs) < 1 {
		return
	}
	s.childs[0].print()
	fmt.Printf(" %s ", s.name)
	if len(s.childs) < 2 {
		return
	}
	s.childs[1].print()
	fmt.Print(")")
}
func (s *SubOpperation) printTree(indentation int) {
	if len(s.childs) < 1 {
		return
	}
	s.childs[0].printTree(indentation + 1)

	printIndentation(indentation)
	fmt.Printf("%s\n", s.name)
	if len(s.childs) < 2 {
		return
	}
	s.childs[1].printTree(indentation + 1)
}

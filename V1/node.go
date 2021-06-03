package V1

import (
	"AlgebraCalculator/V1/log"
	"fmt"
	"math"
)

// Flag indexes that are used as index in flagValues.
const (
	flagNone = 0
	flagRoot = 1

	// Data Types no Children
	flagData     = 10
	flagNumber   = 11
	flagConstant = 12
	flagVariable = 13

	// Action Typs have Children
	flagAction    = 20
	flagOperator1 = 21
	flagOperator2 = 22
	flagVector    = 23
	flagTerm      = 24

	flagBracketRoot = 30
	flagRuleData    = 31

	flagMax      = 40
	flagOptional = 30
)

type node struct {
	childs []*node

	data       string
	dataNumber float64

	flagValues [flagMax]bool
}

func newNode(data string, dataNumber float64, flags ...int) *node {
	node := &node{
		data:       data,
		dataNumber: dataNumber,
	}

	// Set all flagValues that are parsed in.
	for _, flag := range flags {
		node.setFlag(flag, true)
	}

	// If no flagValues are set node will ge flag none
	if len(flags) == 0 {
		node.setFlag(flagNone, true)
	}

	return node
}
func (n *node) setChilds(childs ...*node) {
	n.childs = childs
}
func (n *node) setFlag(flag int, value bool) {
	// Remove none when an other flag is set.
	if n.flagValues[flagNone] && flag != flagNone && value {
		n.flagValues[flagNone] = false
	}

	n.flagValues[flag] = value
}
func (n *node) hasFlag(flags ...int) bool {
	for _, flag := range flags {
		if !n.flagValues[flag] {
			return false
		}
	}
	return true
}
func (n *node) hasAllFlagsOfNode(reference *node) bool {
	for flag, flagValue := range reference.flagValues {
		if flag >= flagOptional {
			break
		}

		if flagValue && !n.hasFlag(flag) {
			return false
		}
	}
	return true
}
func (n *node) equal(reference *node) bool {
	if !n.hasAllFlagsOfNode(reference) {
		return false
	}

	// Number are the same when the number is the same. Else it usually the string.
	if n.hasFlag(flagNumber) {
		return n.dataNumber == reference.dataNumber
	} else {
		return n.data == reference.data
	}
}

// Recursive equal check.
func (n *node) equalDeep(reference *node) bool {
	if len(n.childs) != len(reference.childs) || !n.equal(reference) {
		return false
	}

	for i, child := range n.childs {
		if !child.equal(reference.childs[i]) {
			return false
		}
	}
	return true
}

// Recursive copy
func (n *node) copyDeep() *node {
	copy := newNode(n.data, n.dataNumber)
	copy.flagValues = n.flagValues

	for _, child := range n.childs {
		copy.childs = append(copy.childs, child.copyDeep())
	}
	return copy
}

// Both data types combined create a clear idenity
func (n *node) getIdentity() string { // TODO improve for sorting
	return n.data + fmt.Sprintf("%f", n.dataNumber)
}

// Combined identity of node and childs.
func (n *node) getIdentityDeep() string {
	identify := n.getIdentity()
	for _, child := range n.childs {
		identify += child.getIdentityDeep()
	}
	return identify
}

func (n *node) print() {

	if n.hasFlag(flagBracketRoot) {
		log.Print("(")
	}
	if n.hasFlag(flagVector) {

		log.Print("(")
		for i, child := range n.childs {
			if i != 0 {
				log.Print(", ")
			}
			child.print()
		}
		log.Print(")")

	} else if n.hasFlag(flagOperator2) {
		n.childs[0].print()
		log.Printf(" %s ", n.data)
		n.childs[1].print()
	} else {
		if n.hasFlag(flagNumber) {
			if n.dataNumber == math.Trunc(n.dataNumber) {
				log.Printf("%.0f", n.dataNumber)
			} else {
				log.Printf("%.4f", n.dataNumber)
			}
		} else {
			log.Print(n.data)
		}

		for _, child := range n.childs {
			log.Print(" ")
			child.print()
		}
	}

	if n.hasFlag(flagBracketRoot) {
		log.Print(")")
	}
}
func (n *node) printTree(indentation int) {

	if n.hasFlag(flagOperator2) {
		n.childs[0].printTree(indentation + 1)

		log.Print("\n")
		printIndentation(indentation)
		log.Printf(" %s ", n.data)

		n.childs[1].printTree(indentation + 1)
	} else {

		log.Print("\n")
		printIndentation(indentation)
		if n.hasFlag(flagNumber) {
			if n.dataNumber == math.Trunc(n.dataNumber) {
				log.Printf("%.0f", n.dataNumber)
			} else {
				log.Printf("%.4f", n.dataNumber)
			}
		} else {
			log.Print(n.data)
		}

		for _, child := range n.childs {
			log.Print(" ")
			child.printTree(indentation + 1)
		}
	}
}
func printIndentation(indentation int) {
	for i := 0; i < indentation; i++ {
		if i == indentation-1 {
			log.Print("|> ")
		} else if i == 0 {
			log.Print("|  ")
		} else {
			log.Print("   ")
		}

	}
}

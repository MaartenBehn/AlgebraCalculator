package V4

import (
	"AlgebraCalculator/log"
	"math"
)

const (
	flagNone     = 0
	flagRoot     = 1
	flagRulePart = 2

	// Basic Types no Children
	flagData     = 10
	flagNumber   = 11
	flagConstant = 12
	flagVariable = 13

	// Action Typs have Children
	flagAction    = 20
	flagOperator1 = 21
	flagOperator2 = 22
	flagFraction  = 23

	flagBracketRoot = 30

	flagMax = 40
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
func (n *node) hasFlag(flag int) bool {
	return n.flagValues[flag]
}
func (n *node) hasAllFlagsOfNode(reference *node) bool {
	for flag, flagValue := range reference.flagValues {
		if flagValue && !n.hasFlag(flag) {
			return false
		}
	}
	return true
}
func (n *node) equal(reference *node) bool {
	return n.hasAllFlagsOfNode(reference) && n.data == reference.data && n.dataNumber == reference.dataNumber
}
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
func (n *node) copyDeep() *node {
	copy := newNode(n.data, n.dataNumber)
	copy.flagValues = n.flagValues

	for _, child := range n.childs {
		copy.childs = append(copy.childs, child.copyDeep())
	}
	return copy
}

func (n *node) print() {

	if n.hasFlag(flagBracketRoot) {
		log.Print("(")
	}
	if n.hasFlag(flagOperator2) {
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

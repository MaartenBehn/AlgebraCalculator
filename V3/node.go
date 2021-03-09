package V3

import (
	"AlgebraCalculator/log"
	"strconv"
)

type iNode interface {
	setParent(partent iNode)
	getParent() iNode
	setChilds(childs []iNode)
	getChilds() []iNode
	getMaxChilds() int

	getType() int
	getRank() int
	getDefiner(vaules bool) string
	getDeepDefiner(vaules bool) string
	setBracketRoot(is bool)
	getBracketRoot() bool

	copy() iNode
	solve() bool
	sort() bool
	print()
	printTree(indentation int)
}

const (
	typNone            = 0
	typRoot            = 1
	typVector          = 2
	typVariable        = 3
	typOpperator       = 4
	typMathFunction    = 5
	typSubOperation    = 6
	typTerm            = 7
	typComplexFunction = 8

	rankNone            = 0
	rankRoot            = 1
	rankAppend          = 2
	rankAddSub          = 3
	rankMul             = 4
	rankPow             = 5
	rankMathFunction    = 6
	rankSubOpperation   = 7
	rankTerm            = 8
	rankComplexFunction = 9
	rankNotSolvable     = 100
)

type node struct {
	parent      iNode
	childs      []iNode
	typeId      int
	rank        int
	maxChilds   int
	definer     string
	bracketRoot bool
}

func newNode(typeId int, rank int, maxChilds int) *node {
	return &node{
		typeId:    typeId,
		rank:      rank,
		maxChilds: maxChilds,
		definer:   strconv.Itoa(typeId),
	}
}

func (t *node) setParent(partent iNode) {
	t.parent = partent
}
func (t *node) getParent() iNode {
	return t.parent
}
func (t *node) setChilds(childs []iNode) {
	t.childs = childs
}
func (t *node) getChilds() []iNode {
	return t.childs
}
func (t *node) getMaxChilds() int {
	return t.maxChilds
}

func (t *node) getType() int {
	return t.typeId
}
func (t *node) getRank() int {
	return t.rank
}
func (t *node) getDefiner(vaules bool) string {
	return t.definer
}
func (t *node) getDeepDefiner(vaules bool) string {
	var deepDefiner string
	for _, child := range t.childs {
		deepDefiner += child.getDeepDefiner(vaules)
	}
	deepDefiner += t.definer
	return deepDefiner
}
func (t *node) setBracketRoot(is bool) {
	t.bracketRoot = is
}
func (t *node) getBracketRoot() bool {
	return t.bracketRoot
}

func (t *node) copy() iNode {
	copy := newNode(t.typeId, t.rank, t.maxChilds)
	copy.childs = make([]iNode, len(t.childs))

	for i, child := range t.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (t *node) solve() bool {
	solved := false
	for _, child := range t.childs {
		if child.solve() {
			solved = true
		}
	}
	return solved
}
func (t *node) sort() bool {
	sorted := false
	for _, child := range t.childs {
		if child.sort() {
			sorted = true
		}
	}
	return sorted
}
func (t *node) print() {
	if len(t.childs) > 0 {
		log.Print("(")
		for _, child := range t.childs {
			child.print()
		}
		log.Print(")")
	}
}
func (t *node) printTree(indentation int) {
	log.Print("\n")
	indentation++
	if len(t.childs) > 0 {
		for _, child := range t.childs {
			child.printTree(indentation)
		}
	}
	indentation--
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

// replaceNode replaces old to new and updates the partent and child pointers to new.
func replaceNode(old iNode, new iNode) {

	// Copy node Data to new
	new.setChilds(old.getChilds())
	new.setParent(old.getParent())

	// Set the partents of the childs to new
	for _, child := range new.getChilds() {
		child.setParent(new)
	}

	// Set the childs of the partent to new
	if new.getParent() != nil {
		childs := new.getParent().getChilds()
		for i, child := range childs {
			if child == old {
				childs[i] = new
			}
		}
		new.getParent().setChilds(childs)
	}
}

// insertNode replaces old to new but keep the childs of new while conectiong the partent of old.
func insertNode(old iNode, new iNode) {

	// Copy Partent pointer
	new.setParent(old.getParent())

	// Set new as child of partent of old
	if new.getParent() != nil {
		childs := new.getParent().getChilds()
		for i, child := range childs {
			if child == old {
				childs[i] = new
			}
		}
		new.getParent().setChilds(childs)
	}
}

func pushNode(node iNode, newNode iNode) {
	if node.getType() == typNone {
		replaceNode(node, newNode)
		return
	}

	childs := node.getChilds()
	if len(childs) >= node.getMaxChilds() {
		mostLeftNode := childs[0]
		pushNode(newNode, mostLeftNode)
		childs[0] = newNode
	} else {
		childs = append(childs, newNode)
	}
	node.setChilds(childs)
	newNode.setParent(node)
}

func deepEqual(node1 iNode, node2 iNode) bool {
	return node1.getDeepDefiner(true) == node2.getDeepDefiner(true)
}

package V3

import "fmt"

type INode interface {
	setParent(partent INode)
	getParent() INode
	addChild(child INode, onRightSide bool) INode
	setChilds(childs []INode)
	getChilds() []INode
	getMaxChilds() int

	getType() int
	getRank() int

	copy() INode
	solve()
	simplify()
	print()
}

const (
	TypNone         = 0
	TypVector       = 1
	TypVariable     = 2
	TypOpperator    = 3
	TypFunction     = 4
	TypBrace        = 5
	TypSubOperation = 6
	TypTerm         = 7

	RankNone          = 0
	RankAppend        = 1
	RankAddSub        = 2
	RankMul           = 3
	RankPow           = 4
	RankFunc          = 5
	RankSubOpperation = 6
	RankTerm          = 7
)

type Node struct {
	parent    INode
	childs    []INode
	typeId    int
	rank      int
	maxChilds int
}

func NewNode(typeId int, rank int, maxChilds int) *Node {
	return &Node{
		typeId:    typeId,
		rank:      rank,
		maxChilds: maxChilds,
	}
}

func (t *Node) setParent(partent INode) {
	t.parent = partent
}
func (t *Node) getParent() INode {
	return t.parent
}
func (t *Node) addChild(child INode, onRightSide bool) INode {
	if onRightSide {
		t.childs = append(t.childs, child)
	} else {
		t.childs = append([]INode{child}, t.childs...)
	}
	return child
}
func (t *Node) setChilds(childs []INode) {
	t.childs = childs
}
func (t *Node) getChilds() []INode {
	return t.childs
}
func (t *Node) getMaxChilds() int {
	return t.maxChilds
}

func (t Node) getType() int {
	return t.typeId
}
func (t Node) getRank() int {
	return t.rank
}

func (t *Node) copy() INode {
	copy := NewNode(t.typeId, t.rank, t.maxChilds)
	copy.childs = make([]INode, len(t.childs))

	for i, child := range t.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (t *Node) solve() {
	for _, child := range t.childs {
		child.solve()
	}
}
func (t *Node) simplify() {
	for _, child := range t.childs {
		child.simplify()
	}
}
func (t *Node) print() {
	if len(t.childs) > 0 {
		fmt.Print("(")
		for _, child := range t.childs {
			child.print()
		}
		fmt.Print(")")
	}
}

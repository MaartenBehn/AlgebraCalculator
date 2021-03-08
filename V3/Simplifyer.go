package V3

import "strconv"

const (
	TypSimpVec = 1001
	TypSimpVar = 1002
	TypSimpVer = 1003
	TypSimpAll = 1004
)

type SimpNode struct {
	*Node
	id       int
	allIds   bool
	inverted bool
}

func NewSimpNode(typeId int, id int, inverted bool, allIds bool) *SimpNode {
	node := &SimpNode{
		Node:     NewNode(typeId, RankNotSolvable, 0),
		id:       id,
		inverted: inverted,
		allIds:   allIds,
	}
	node.definer += strconv.Itoa(id)
	return node
}

func (s *SimpNode) copy() INode {
	return NewSimpNode(s.typeId, s.id, s.inverted, s.allIds)
}

type SimpDataBuffer struct {
	nodes []INode
}

func (s *SimpDataBuffer) checkSimpNode(node INode, simpNode INode) bool {

	index := len(s.nodes)
	for i, testnode := range s.nodes {
		if deepEqual(node, testnode) {
			index = i
			break
		}
	}

	if simpNode.(*SimpNode).id >= len(s.nodes) {
		s.nodes = append(s.nodes, node)
	}

	debug := simpNode.(*SimpNode).id == index || (simpNode.(*SimpNode).allIds && simpNode.(*SimpNode).id >= index)
	return debug
}

func (s *SimpDataBuffer) getReplacement(node INode) INode {
	replacement := node

	if node.getType() >= 1000 {
		replacement = s.nodes[node.(*SimpNode).id]
	}

	return replacement
}

type SimpRule struct {
	line    int
	base    string
	search  INode
	replace INode
}

func (s SimpRule) applyRule(node INode, simpNode INode) bool {
	for _, child := range node.getChilds() {
		if s.applyRule(child, simpNode) {
			return true
		}
	}

	dataBuffer := &SimpDataBuffer{}
	if simpDoesMap(node, s.search, dataBuffer) {

		replacement := dataBuffer.getReplacement(s.replace.copy())

		insertNode(node, replacement)

		for _, child := range replacement.getChilds() {
			s.simpInsert(child, dataBuffer)
		}
		return true
	}
	return false
}

func simpDoesMap(node INode, simpNode INode, dataBuffer *SimpDataBuffer) bool {
	for i, child := range simpNode.getChilds() {

		if i >= len(node.getChilds()) {
			return false
		}

		if !simpDoesMap(node.getChilds()[i], child, dataBuffer) {
			return false
		}
	}

	switch simpNode.getType() {
	case TypSimpVec:
		return (node.getType() == TypVector || (node.getType() != TypVector && simpNode.(*SimpNode).inverted)) &&
			dataBuffer.checkSimpNode(node, simpNode)

	case TypSimpVar:
		return (node.getType() == TypVariable || (node.getType() != TypVariable && simpNode.(*SimpNode).inverted)) &&
			dataBuffer.checkSimpNode(node, simpNode)

	case TypSimpAll:
		return dataBuffer.checkSimpNode(node, simpNode) && !simpNode.(*SimpNode).inverted

	case TypOpperator:
		return node.getType() == TypOpperator && node.(INamedNode).getName() == simpNode.(INamedNode).getName()

	case TypMathFunction:
		return node.getType() == TypMathFunction && node.(INamedNode).getName() == simpNode.(INamedNode).getName()

	case TypVector:
		return node.getType() == TypVector && len(node.(*Vector).values) == 1 && node.(*Vector).values[0] == simpNode.(*Vector).values[0]

	}
	return false
}

func (s SimpRule) simpInsert(simpNode INode, dataBuffer *SimpDataBuffer) {

	replacement := dataBuffer.getReplacement(simpNode)

	insertNode(simpNode, replacement)

	for _, child := range replacement.getChilds() {
		s.simpInsert(child, dataBuffer)
	}
}

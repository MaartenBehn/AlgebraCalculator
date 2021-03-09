package V3

import "strconv"

const (
	typSimpVec = 1001
	typSimpVar = 1002
	typSimpVer = 1003
	typSimpAll = 1004
)

type simpNode struct {
	*node
	id       int
	allIds   bool
	inverted bool
}

func newSimpNode(typeId int, id int, inverted bool, allIds bool) *simpNode {
	node := &simpNode{
		node:     newNode(typeId, rankNotSolvable, 0),
		id:       id,
		inverted: inverted,
		allIds:   allIds,
	}
	node.definer += strconv.Itoa(id)
	return node
}

func (s *simpNode) copy() iNode {
	return newSimpNode(s.typeId, s.id, s.inverted, s.allIds)
}

type simpDataBuffer struct {
	nodes []iNode
}

func (s *simpDataBuffer) checkSimpNode(node iNode, nodeSimp iNode) bool {

	index := len(s.nodes)
	for i, testnode := range s.nodes {
		if deepEqual(node, testnode) {
			index = i
			break
		}
	}

	if nodeSimp.(*simpNode).id >= len(s.nodes) {
		s.nodes = append(s.nodes, node)
	}

	debug := nodeSimp.(*simpNode).id == index || (nodeSimp.(*simpNode).allIds && nodeSimp.(*simpNode).id >= index)
	return debug
}
func (s *simpDataBuffer) getReplacement(node iNode) iNode {
	replacement := node

	if node.getType() >= 1000 {
		replacement = s.nodes[node.(*simpNode).id]
	}

	return replacement
}

type simpRule struct {
	line    int
	base    string
	search  iNode
	replace iNode
}

func (s simpRule) applyRule(node iNode, simpNode iNode) bool {
	for _, child := range node.getChilds() {
		if s.applyRule(child, simpNode) {
			return true
		}
	}

	dataBuffer := &simpDataBuffer{}
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
func simpDoesMap(node iNode, nodeSimp iNode, dataBuffer *simpDataBuffer) bool {
	for i, child := range nodeSimp.getChilds() {

		if i >= len(node.getChilds()) {
			return false
		}

		if !simpDoesMap(node.getChilds()[i], child, dataBuffer) {
			return false
		}
	}

	switch nodeSimp.getType() {
	case typSimpVec:
		return (node.getType() == typVector || (node.getType() != typVector && nodeSimp.(*simpNode).inverted)) &&
			dataBuffer.checkSimpNode(node, nodeSimp)

	case typSimpVar:
		return (node.getType() == typVariable || (node.getType() != typVariable && nodeSimp.(*simpNode).inverted)) &&
			dataBuffer.checkSimpNode(node, nodeSimp)

	case typSimpAll:
		return dataBuffer.checkSimpNode(node, nodeSimp) && !nodeSimp.(*simpNode).inverted

	case typOpperator:
		return node.getType() == typOpperator && node.(iNamedNode).getName() == nodeSimp.(iNamedNode).getName()

	case typMathFunction:
		return node.getType() == typMathFunction && node.(iNamedNode).getName() == nodeSimp.(iNamedNode).getName()

	case typVector:
		return node.getType() == typVector && len(node.(*vector).values) == 1 && node.(*vector).values[0] == nodeSimp.(*vector).values[0]

	}
	return false
}
func (s simpRule) simpInsert(simpNode iNode, dataBuffer *simpDataBuffer) {

	replacement := dataBuffer.getReplacement(simpNode)

	insertNode(simpNode, replacement)

	for _, child := range replacement.getChilds() {
		s.simpInsert(child, dataBuffer)
	}
}

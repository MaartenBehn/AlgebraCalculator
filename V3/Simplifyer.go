package V3

const (
	TypSimpVector   = 1001
	TypSimpVariable = 1002
)

type SimpNode struct {
	*Node
	id int
}

func NewSimpNode(typeId int, id int) *SimpNode {
	return &SimpNode{
		Node: NewNode(typeId, 0, 0),
		id:   id,
	}
}

type SimpIdMapper struct {
	vectors       []*Vector
	sameVectors   []int
	variables     []*Variable
	sameVariables []int
}

func (s *SimpIdMapper) addVectorId(vector *Vector) int {
	s.vectors = append(s.vectors, vector)
	for i, testVector := range s.vectors {
		if testVector == vector { // TODO: Check if this actually works

			return i
		}
	}
	return len(s.vectors) - 1
}

func (s *SimpIdMapper) addVariableId(variable *Variable) int {
	s.variables = append(s.variables, variable)
	for i, testVariable := range s.variables {
		if testVariable.name == variable.name {
			return i
		}
	}
	return len(s.variables) - 1
}

var simpRules []SimpRule

type SimpRule struct {
	search  INode
	replace INode
}

func (s SimpRule) tryRule(node INode, simpNode INode) bool {
	for _, child := range node.getChilds() {
		if s.tryRule(child, simpNode) {
			return true
		}
	}

	idMapper := &SimpIdMapper{}
	if simpEqual(node, s.search, idMapper) {

		s.simpReplace(node, s.replace, idMapper)

		return true
	}
	return false
}

func simpEqual(node INode, simpNode INode, mapper *SimpIdMapper) bool {
	for i, child := range simpNode.getChilds() {

		if i >= len(node.getChilds()) {
			return false
		}

		if !simpEqual(node.getChilds()[i], child, mapper) {
			return false
		}
	}

	switch simpNode.getType() {
	case TypSimpVector:
		return node.getType() == TypVector && mapper.addVectorId(node.(*Vector)) <= simpNode.(*SimpNode).id

	case TypSimpVariable:
		return node.getType() == TypVariable && mapper.addVariableId(node.(*Variable)) <= simpNode.(*SimpNode).id

	case TypOpperator:
		return node.getType() == TypOpperator && node.(INamedNode).getName() == simpNode.(INamedNode).getName()

	case TypFunction:
		return node.getType() == TypFunction && node.(INamedNode).getName() == simpNode.(INamedNode).getName()

	case TypVector:
		return node.getType() == TypVector && len(node.(*Vector).values) == 1 && node.(*Vector).values[0] == simpNode.(*Vector).values[0]

	}

	return false
}

func (s SimpRule) simpReplace(node INode, simpNode INode, mapper *SimpIdMapper) {
	node.setChilds(nil)
	for _, simpChild := range simpNode.getChilds() {
		child := NewNode(TypNone, RankNone, 0)
		puschChild(child, node)
		s.simpReplace(child, simpChild, mapper)
	}

	replace := simpNode.copy()
	switch simpNode.getType() {
	case TypSimpVector:
		replace = mapper.vectors[simpNode.(*SimpNode).id].copy()

	case TypSimpVariable:
		replace = mapper.variables[simpNode.(*SimpNode).id].copy()
	}

	replaceNode(node, replace)
}

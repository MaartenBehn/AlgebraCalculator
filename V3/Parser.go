package V3

import (
	"strings"
)

type ParseData struct {
	root      INode
	current   INode
	braceEnd  INode
	found     bool
	err       error
	variables []*Variable
	terms     []INode
}

var termParseFunc = []func(part string, data *ParseData){
	tryParseSolvableNodes,
	tryParseVariables,
	tryParseTerm,
	tryParseNumber,
}

func parseTerm(text string, terms []INode) (*Term, error) {

	parts := removeEmptiStrings(strings.Split(text, "="))

	if len(parts) != 2 {
		return nil, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "There is not a valid part before or after \"=\" in term!")
	}
	parts1 := removeEmptiStrings(splitAny(parts[0], " <>"))

	var variables []*Variable
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, NewVariable(parts1[i]))
	}
	parts2 := removeEmptiStrings(splitAny(parts[1], " <>"))

	term := NewTerm(parts1[0], variables)
	term.root = NewNode(TypRoot, RankRoot, 1)

	root, err, _ := parseRoot(parts2, termParseFunc, variables, terms)
	if handelError(err) {
		return &Term{}, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Term could not be parsed!")
	}
	term.root.setChilds([]INode{root})
	root.setParent(term.root)

	return term, nil
}

var simpRuleParseFunc = []func(part string, data *ParseData){
	tryParseRuleSymbols,
	tryParseSolvableNodes,
	tryParseNumber,
}

func parseSimpRule(text string) (SimpRule, error) {
	parts := removeEmptiStrings(strings.Split(text, "="))

	if len(parts) != 2 {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelFatal, "Rule has not excatily one = !")
	}

	parts1 := strings.Split(parts[0], " ")
	root1, err, _ := parseRoot(parts1, simpRuleParseFunc, nil, nil)
	if handelError(err) {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root1.getType() == TypNone {
		root1 = root1.getChilds()[0]
		root1.setParent(nil)
	}

	parts2 := strings.Split(parts[1], " ")
	root2, err, _ := parseRoot(parts2, simpRuleParseFunc, nil, nil)
	if handelError(err) {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root2.getType() == TypNone {
		root2 = root2.getChilds()[0]
		root2.setParent(nil)
	}

	return SimpRule{
		base:    text,
		search:  root1,
		replace: root2,
	}, nil
}

func tryParseRuleSymbols(part string, data *ParseData) {
	if strings.Contains(part, "_") {

		inverted := strings.Contains(part, "!")
		invertedOffset := 0
		if inverted {
			invertedOffset = 1
		}

		typ := 0
		switch part[:3+invertedOffset] {
		case "vec":
			typ = TypSimpVec
			break
		case "var":
			typ = TypSimpVar
			break
		case "ver":
			typ = TypSimpVer
			break
		case "all":
			typ = TypSimpAll
			break
		}
		id := getInt(part[4+invertedOffset])
		allIds := len(part) >= (6+invertedOffset) && part[5+invertedOffset] == 'i'

		simpNode := NewSimpNode(typ, id, inverted, allIds)
		addParsedNode(simpNode, data)
		data.found = true
	}
}
func tryParseSolvableNodes(part string, data *ParseData) {
	for _, solvableTermNode := range solvableTermNodes {
		if part == solvableTermNode.getName() {
			termNode := solvableTermNode.copy()

			termNode = addParsedNode(termNode, data)
			data.current = termNode
			data.found = true
			break
		}
	}
}
func tryParseVariables(part string, data *ParseData) {
	for _, variable := range data.variables {
		if variable.getName() == part {
			addParsedNode(variable.copy(), data)
			data.found = true
		}
	}
}
func tryParseTerm(part string, data *ParseData) {
	for _, testTerm := range data.terms {
		if part == testTerm.(*Term).name {
			copy := testTerm.copy()
			copy.setBracketRoot(true)
			copy = addParsedNode(copy, data)
			data.current = copy
			data.found = true
		}
	}
}
func tryParseNumber(part string, data *ParseData) {
	if isNumber(part) {
		var vector *Vector
		vector, data.err = getVector(part)
		addParsedNode(vector, data)
		data.found = true
	}

	if strings.Contains(part, ".") {
		parts2 := removeEmptiStrings(strings.Split(part, "."))
		if len(parts2) != 2 {
			return
		}

		if isNumber(parts2[0]) {
			var vector *Vector
			vector, data.err = getVector(parts2[1])
			addParsedNode(vector, data)
			data.found = true
		}
	}
}

func parseRoot(parts []string, funcs []func(part string, data *ParseData), varibles []*Variable, terms []INode) (INode, error, int) {
	parts = removeEmptiStrings(parts)
	data := &ParseData{
		root:      NewNode(TypNone, RankNone, 0),
		variables: varibles,
		terms:     terms,
	}
	data.current = data.root

	var i int
	for i = 0; i < len(parts); i++ {
		part := parts[i]
		data.found = false
		if part == "(" {
			subRoot, err, index := parseRoot(parts[i+1:], funcs, varibles, terms)
			data.err = err
			subRoot.setBracketRoot(true)
			data.current.setChilds(append(data.current.getChilds(), subRoot))
			subRoot.setParent(data.current)

			i += index + 1
			data.found = true
		}

		if part == ")" {
			break
		}

		for _, function := range funcs {
			function(part, data)
		}

		if !data.found {
			data.err = NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Expression part "+part+" could not be parsed!")
		}

		if data.err != nil {
			return nil, data.err, 0
		}
	}

	if data.root.getType() == TypNone && len(data.root.getChilds()) == 1 {
		data.root = data.root.getChilds()[0]
	}

	return data.root, nil, i
}

func addParsedNode(node INode, data *ParseData) INode {
	current := data.current

	if current.getType() == TypNone {

		replaceNode(current, node)

	} else if node.getRank() > current.getRank() || node.getBracketRoot() {

		childs := current.getChilds()
		if len(childs) < current.getMaxChilds() {
			childs = append(childs, node)
		} else {
			child := childs[len(childs)-1]
			childs[len(childs)-1] = node
			node.setChilds([]INode{child})
			child.setParent(node)
		}
		current.setChilds(childs)
		node.setParent(current)

	} else if node.getRank() <= current.getRank() {

		if current.getParent() == nil {
			current.setParent(node)
			node.setChilds([]INode{current})
		} else {
			data.current = current.getParent()
			node = addParsedNode(node, data)
		}
	}

	if node.getParent() == nil {
		data.root = node
		data.current = node
	}

	return node
}

package AlgebraCalculator

import (
	"strings"
)

type parseData struct {
	root      iNode
	current   iNode
	braceEnd  iNode
	found     bool
	err       error
	variables []*variable
	terms     map[int]iNode
}

var termParseFunc = []func(part string, data *parseData){
	tryParseSolvableNodes,
	tryParseVariables,
	tryParseTerm,
	tryParseNumber,
}

func parseTerm(text string, terms map[int]iNode) (*term, error) {

	parts := removeEmptiStrings(strings.Split(text, "="))

	if len(parts) != 2 {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "There is not a valid part before or after \"=\" in term!")
	}
	parts1 := removeEmptiStrings(splitAny(parts[0], " <>"))

	var variables []*variable
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, newVariable(parts1[i]))
	}
	parts2 := removeEmptiStrings(splitAny(parts[1], " <>"))

	t := newTerm(parts1[0], variables)
	t.root = newNode(typRoot, rankRoot, 1)

	root, err, _ := parseRoot(parts2, termParseFunc, variables, terms)
	if handelError(err) {
		return &term{}, newError(errorTypParsing, errorCriticalLevelPartial, "term could not be parsed!")
	}
	t.root.setChilds([]iNode{root})
	root.setParent(t.root)

	return t, nil
}

var simpRuleParseFunc = []func(part string, data *parseData){
	tryParseRuleSymbols,
	tryParseSolvableNodes,
	tryParseNumber,
}

func parseSimpRule(text string) (simpRule, error) {
	parts := removeEmptiStrings(strings.Split(text, "="))

	if len(parts) != 2 {
		return simpRule{}, newError(errorTypParsing, errorCriticalLevelFatal, "Rule has not excatily one = !")
	}

	parts1 := strings.Split(parts[0], " ")
	root1, err, _ := parseRoot(parts1, simpRuleParseFunc, nil, nil)
	if handelError(err) {
		return simpRule{}, newError(errorTypParsing, errorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root1.getType() == typNone {
		root1 = root1.getChilds()[0]
		root1.setParent(nil)
	}

	parts2 := strings.Split(parts[1], " ")
	root2, err, _ := parseRoot(parts2, simpRuleParseFunc, nil, nil)
	if handelError(err) {
		return simpRule{}, newError(errorTypParsing, errorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root2.getType() == typNone {
		root2 = root2.getChilds()[0]
		root2.setParent(nil)
	}

	return simpRule{
		base:    text,
		search:  root1,
		replace: root2,
	}, nil
}

func tryParseRuleSymbols(part string, data *parseData) {
	if strings.Contains(part, "_") {

		inverted := strings.Contains(part, "!")
		invertedOffset := 0
		if inverted {
			invertedOffset = 1
		}

		typ := 0
		switch part[:3+invertedOffset] {
		case "vec":
			typ = typSimpVec
			break
		case "var":
			typ = typSimpVar
			break
		case "ver":
			typ = typSimpVer
			break
		case "all":
			typ = typSimpAll
			break
		}
		id := getInt(part[4+invertedOffset])
		allIds := len(part) >= (6+invertedOffset) && part[5+invertedOffset] == 'i'

		simpNode := newSimpNode(typ, id, inverted, allIds)
		addParsedNode(simpNode, data)
		data.found = true
	}
}
func tryParseSolvableNodes(part string, data *parseData) {
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
func tryParseVariables(part string, data *parseData) {
	for _, variable := range data.variables {
		if variable.getName() == part {
			addParsedNode(variable.copy(), data)
			data.found = true
		}
	}
}
func tryParseTerm(part string, data *parseData) {
	for i := len(data.terms) - 1; i > 0; i-- {
		testTerm := data.terms[i]
		if part == testTerm.(*term).name {
			copy := testTerm.copy()
			copy.setBracketRoot(true)
			copy = addParsedNode(copy, data)
			data.current = copy
			data.found = true
			break
		}
	}
}
func tryParseNumber(part string, data *parseData) {
	if isNumber(part) {
		var vector *vector
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
			var vector *vector
			vector, data.err = getVector(parts2[1])
			addParsedNode(vector, data)
			data.found = true
		}
	}
}

func parseRoot(parts []string, funcs []func(part string, data *parseData), varibles []*variable, terms map[int]iNode) (iNode, error, int) {
	parts = removeEmptiStrings(parts)
	data := &parseData{
		root:      newNode(typNone, rankNone, 0),
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
			data.err = newError(errorTypParsing, errorCriticalLevelPartial, "Expression part "+part+" could not be parsed!")
		}

		if data.err != nil {
			return nil, data.err, 0
		}
	}

	if data.root.getType() == typNone && len(data.root.getChilds()) == 1 {
		data.root = data.root.getChilds()[0]
	}

	if data.root.getBracketRoot() {
		data.root.setBracketRoot(false)
	}

	return data.root, nil, i
}

func addParsedNode(node iNode, data *parseData) iNode {
	current := data.current

	if current.getType() == typNone {

		replaceNode(current, node)

	} else if node.getRank() > current.getRank() || node.getBracketRoot() {

		childs := current.getChilds()
		if len(childs) < current.getMaxChilds() {
			childs = append(childs, node)
		} else {
			child := childs[len(childs)-1]
			childs[len(childs)-1] = node
			node.setChilds([]iNode{child})
			child.setParent(node)
		}
		current.setChilds(childs)
		node.setParent(current)

	} else if node.getRank() <= current.getRank() {

		if current.getParent() == nil {
			current.setParent(node)
			node.setChilds([]iNode{current})
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

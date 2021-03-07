package V3

import (
	"strings"
)

type ParseData struct {
	root     INode
	current  INode
	braceEnd INode
	found    bool
	err      error
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

	parseData := &ParseData{
		root: NewNode(TypNone, RankNone, 0),
	}
	parseData.current = parseData.root
	for _, part := range parts2 {
		tryParseBracket(part, parseData)
		tryParseSolvableNodes(part, parseData)
		tryParseVariables(part, variables, parseData)
		tryParseTerm(part, terms, parseData)
		tryParseNumber(part, parseData)

		if !parseData.found {
			parseData.err = NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Expression part "+part+" could not be parsed!")
		}

		if handelError(parseData.err) {
			return nil, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Term could not be parsed!")
		}
	}

	term := NewTerm(parts1[0], variables)
	term.root = parseData.root

	if term.root.getType() != TypNone {
		root := NewNode(TypNone, RankNone, 0)
		term.root.setParent(root)
		root.addChild(term.root, true)
		term.root = root
	}

	return term, nil
}

func parseSimpRule(text string) (SimpRule, error) {
	parts := removeEmptiStrings(strings.Split(text, "="))

	if len(parts) != 2 {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelFatal, "Rule has not excatily one = !")
	}

	parts1 := strings.Split(parts[0], " ")
	root1, err := parseSimpRuleRoot(parts1)
	if handelError(err) {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root1.getType() == TypNone {
		root1 = root1.getChilds()[0]
		root1.setParent(nil)
	}

	parts2 := strings.Split(parts[1], " ")
	root2, err := parseSimpRuleRoot(parts2)
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
func parseSimpRuleRoot(parts []string) (INode, error) {
	parts = removeEmptiStrings(parts)
	parseData := &ParseData{
		root: NewNode(TypNone, RankNone, 0),
	}
	parseData.current = parseData.root
	for _, part := range parts {
		tryParseBracket(part, parseData)
		tryParseRuleSymbols(part, parseData)
		tryParseSolvableNodes(part, parseData)
		tryParseNumber(part, parseData)

		if !parseData.found {
			parseData.err = NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Expression part "+part+" could not be parsed!")
		}

		if parseData.err != nil {
			return nil, parseData.err
		}
	}

	if parseData.root.getType() == TypNone && len(parseData.root.getChilds()) == 1 {
		parseData.root = parseData.root.getChilds()[0]
	}

	return parseData.root, nil
}

func tryParseBracket(part string, data *ParseData) {
	if part == "(" {
		data.braceEnd = data.current

		data.err = puschChild(NewNode(TypNone, RankNone, 0), data.current)
		data.current = data.current.getChilds()[len(data.current.getChilds())-1]
		data.found = true
	}

	if part == ")" {
		data.current = data.braceEnd
		data.found = true
	}
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
		data.err = puschChild(simpNode, data.current)
		data.found = true
	}
}
func tryParseSolvableNodes(part string, data *ParseData) {
	for _, solvableTermNode := range solvableTermNodes {
		if part == solvableTermNode.getName() {
			termNode := solvableTermNode.copy()
			rank := termNode.getRank()

			parentNode := spotAcordingToRank(data.current, rank)

			if parentNode.getType() == TypNone {
				if data.braceEnd != nil && parentNode == data.braceEnd {

					childs := parentNode.getChilds()
					termNode.setParent(parentNode)
					parentNode.setChilds([]INode{termNode})
					for _, child := range childs {
						child.setParent(termNode)
					}
					termNode.setChilds(childs)

					data.current = termNode
				} else {
					replaceNode(parentNode, termNode)
					data.current = termNode

					if data.current.getParent() == nil {
						data.root = data.current
					}
				}

			} else {
				data.err = puschChild(termNode, parentNode)
				data.current = termNode
			}
			data.found = true
			break
		}
	}
}
func tryParseVariables(part string, variables []*Variable, data *ParseData) {
	for _, variable := range variables {
		if variable.getName() == part {
			data.err = puschChild(variable.copy(), data.current)
			data.found = true
		}
	}
}
func tryParseTerm(part string, terms []INode, data *ParseData) {
	for _, testTerm := range terms {
		if part == testTerm.(*Term).name {
			copy := testTerm.copy()
			data.err = puschChild(copy, data.current)
			data.current = copy
			data.found = true
		}
	}
}
func tryParseNumber(part string, data *ParseData) {
	if isNumber(part) {
		var vector *Vector
		vector, data.err = getVector(part)
		data.err = puschChild(vector, data.current)
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
			data.err = puschChild(vector, data.current)
			data.found = true
		}
	}
}

func spotAcordingToRank(current INode, rank int) INode {
	if current.getRank() >= rank {
		partent := current.getParent()

		if partent == nil {
			partent = NewNode(0, 0, 0)
			current.setParent(partent)
			partent.setChilds([]INode{current})
		}

		if partent.getRank() <= current.getRank() {
			return spotAcordingToRank(partent, rank)
		}

		return current
	} else {
		return current
	}
}
func puschChild(child INode, inNode INode) error {

	childs := inNode.getChilds()
	if inNode.getType() != 0 && inNode.getMaxChilds() == 0 {
		return NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Can't push Child becuase maxChilds is zero!")
	} else if inNode.getType() != 0 && len(childs) > 0 && len(childs) >= inNode.getMaxChilds() {
		mostRightChild := childs[len(childs)-1]
		err := puschChild(mostRightChild, child)
		if err != nil {
			return err
		}
		childs[len(childs)-1] = child
		inNode.setChilds(childs)
	} else {
		inNode.addChild(child, true)
	}
	child.setParent(inNode)
	return nil
}

package V3

import (
	"strings"
)

func parseTerm(text string) (*Term, error) {

	parts := strings.Split(text, "=")
	parts = removeEmptiStrings(parts)

	if len(parts) != 2 {
		return nil, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "There is not a valid part before or after \"=\" in term!")
	}
	parts1 := removeEmptiStrings(splitAny(parts[0], " <>"))

	var variables []*Variable
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, NewVariable(parts1[i]))
	}
	parts2 := splitAny(parts[1], " <>")

	root, err := parseRoot(parts2, variables)
	if handelError(err) {
		return nil, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Term could not be parsed!")
	}

	term := NewTerm(parts1[0], variables)
	err = puschChild(root, term)
	if err != err {
		panic(err)
	}

	return term, nil
}

func parseRoot(parts []string, variables []*Variable) (INode, error) {
	parts = removeEmptiStrings(parts)

	var currentNode INode = NewNode(TypNone, RankNone, 0)
	root := currentNode

	var barceEnd INode
	var err error
	for _, part := range parts {

		found := false

		if part == "(" {
			barceEnd = currentNode

			err = puschChild(NewNode(TypNone, RankNone, 0), currentNode)
			currentNode = currentNode.getChilds()[len(currentNode.getChilds())-1]
			found = true
		}

		if part == ")" {
			currentNode = barceEnd
			found = true
		}

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
			err = puschChild(simpNode, currentNode)
			found = true
		}

		for _, solvableTermNode := range solvableTermNodes {
			if part == solvableTermNode.getName() {
				termNode := solvableTermNode.copy()
				rank := termNode.getRank()

				parentNode := spotAcordingToRank(currentNode, rank)

				if parentNode.getType() == TypNone {
					replaceNode(parentNode, termNode)
					currentNode = termNode

					if currentNode.getParent() == nil {
						root = currentNode
					}

				} else {
					err = puschChild(termNode, parentNode)
					currentNode = termNode
				}
				found = true
				break
			}
		}

		for _, variable := range variables {
			if variable.getName() == part {
				err = puschChild(variable.copy(), currentNode)
				found = true
			}
		}

		if strings.Contains(part, ".") {
			parts2 := strings.Split(part, ".")
			parts2 = removeEmptiStrings(parts2)

			if len(parts2) != 2 {

				found = true
			}

			if parts2[0] == ")" {

				found = true
			}

			if isNumber(parts2[0]) {

				found = true
			}
		}

		if isNumber(part) {
			err = puschChild(getVector(part), currentNode)
			found = true
		}

		if !found {
			err = NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Expression part "+part+" could not be parsed!")
		}

		if err != nil {
			return nil, err
		}
	}
	return root, nil
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

func parseSimpRule(text string) (SimpRule, error) {
	parts := strings.Split(text, "=")

	if len(parts) != 2 {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelFatal, "Rule has not excatily one = !")
	}

	parts1 := strings.Split(parts[0], " ")
	root1, err := parseRoot(parts1, nil)
	if handelError(err) {
		return SimpRule{}, NewError(ErrorTypParsing, ErrorCriticalLevelPartial, "Rule could not be parsed!")
	}

	if root1.getType() == TypNone {
		root1 = root1.getChilds()[0]
		root1.setParent(nil)
	}

	parts2 := strings.Split(parts[1], " ")
	root2, err := parseRoot(parts2, nil)
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

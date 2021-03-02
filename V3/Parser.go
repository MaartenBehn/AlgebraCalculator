package V3

import (
	"log"
	"strings"
)

func parseTerm(text string) *Term {

	parts := strings.Split(text, "=")

	if len(parts) != 2 {
		log.Panic("Invalid Variable creation!")
	}
	parts1 := removeEmptiStrings(splitAny(parts[0], " <>"))

	var variables []*Variable
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, NewVariable(parts1[i]))
	}
	parts2 := splitAny(parts[1], " <>")

	root := parseRoot(parts2, variables)

	term := NewTerm(parts1[0], variables)
	puschChild(root, term)

	return term
}

func parseRoot(parts []string, variables []*Variable) INode {
	parts = removeEmptiStrings(parts)

	var currentNode INode = NewNode(TypNone, RankNone, 0)
	root := currentNode
	for _, part := range parts {

		if part == "(" {
			puschChild(NewNode(TypNone, RankNone, 0), currentNode)
			currentNode = currentNode.getChilds()[len(currentNode.getChilds())-1]
			continue
		}

		if part == ")" {
			currentNode = currentNode.getParent()
			continue
		}

		if strings.Contains(part, "vec_") {

			id := getInt(part[4])
			puschChild(NewSimpNode(TypSimpVector, id), currentNode)
		}

		if strings.Contains(part, "var_") {

			id := getInt(part[4])
			puschChild(NewSimpNode(TypSimpVariable, id), currentNode)
		}

		for _, solvableTermNode := range solvableTermNodes {
			if part == solvableTermNode.getName() {
				termNode := solvableTermNode.copy()
				rank := termNode.getRank()

				parentNode, parentRank := findEmptiInPartentWithRankLevel(currentNode, rank)

				if parentRank >= rank {

					replaceNode(parentNode, termNode)
					currentNode = termNode

					if currentNode.getParent() == nil {
						root = currentNode
					}
				} else {
					puschChild(termNode, currentNode)
					currentNode = termNode
				}
			}
		}

		for _, variable := range variables {
			if variable.getName() == part {
				puschChild(variable.copy(), currentNode)
			}
		}

		if strings.Contains(part, ".") {
			parts2 := strings.Split(part, ".")
			parts2 = removeEmptiStrings(parts2)

			if len(parts2) != 2 {
				continue
			}

			if parts2[0] == ")" {

				continue
			}

			if isNumber(parts2[0]) {

				continue
			}
		}

		if isNumber(part) {

			puschChild(getVector(part), currentNode)
			continue
		}
	}
	return root
}

func findEmptiInPartentWithRankLevel(current INode, rankLevel int) (INode, int) {
	if current.getType() == TypNone {
		return current, rankLevel
	} else if current.getRank() < rankLevel {
		rankLevel = current.getRank()
	}

	if current.getParent() != nil {
		return findEmptiInPartentWithRankLevel(current.getParent(), rankLevel)
	}

	newPartent := &Node{}
	current.setParent(newPartent)
	newPartent.addChild(current, true)
	return findEmptiInPartentWithRankLevel(newPartent, rankLevel)
}

func puschChild(child INode, inNode INode) {

	childs := inNode.getChilds()
	if inNode.getType() != 0 && inNode.getMaxChilds() == 0 {
		panic("Cant add Child!")
	} else if inNode.getType() != 0 && len(childs) > 0 && len(childs) >= inNode.getMaxChilds() {
		mostRightChild := childs[len(childs)-1]
		puschChild(mostRightChild, child)
		childs[len(childs)-1] = child
		inNode.setChilds(childs)
	} else {
		inNode.addChild(child, true)
	}
	child.setParent(inNode)
}
func replaceNode(old INode, new INode) {

	new.setChilds(old.getChilds())
	new.setParent(old.getParent())

	for _, child := range new.getChilds() {
		child.setParent(new)
	}

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

func parseSimpRule(text string) SimpRule {
	parts := strings.Split(text, "=")

	if len(parts) != 2 {
		log.Panic("Invalid Rule creation!")
	}

	parts1 := strings.Split(parts[0], " ")
	root1 := parseRoot(parts1, nil)

	if root1.getType() == TypNone {
		root1 = root1.getChilds()[0]
		root1.setParent(nil)
	}

	parts2 := strings.Split(parts[1], " ")
	root2 := parseRoot(parts2, nil)

	if root2.getType() == TypNone {
		root2 = root2.getChilds()[0]
		root2.setParent(nil)
	}

	return SimpRule{
		search:  root1,
		replace: root2,
	}
}

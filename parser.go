package AlgebraCalculator

import (
	"strconv"
)

const (
	rankNone         = 0
	rankRoot         = 1
	rankAppend       = 2
	rankAddSub       = 3
	rankMul          = 4
	rankPow          = 5
	rankMathFunction = 6
	rankSubOperation = 7
	rankTermFunction = 8
	rankTerm         = 9
	rankTermEnd      = 10
)

func initParser() {
	parseTermFuncs = append(parseTermFuncs,
		func(text string) *parserNode { return tryParseNumber(text) },

		func(text string) *parserNode { return tryParseOperator2(text, "+", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "-", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "*", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "/", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "pow", rankPow) },

		func(text string) *parserNode { return tryParseOperator2(text, ",", rankAppend) },
		func(text string) *parserNode { return tryParseOperator2(text, ".", rankSubOperation) },

		func(text string) *parserNode { return tryParseOperator1(text, "sin", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "sinh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "asin", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "asinh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "cos", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "cosh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "acos", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "acosh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "tan", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "tanh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "atan", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "atanh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator2(text, "atan2", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "abs", rankMathFunction) },

		func(text string) *parserNode { return tryParseOperator2(text, "dot", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "len", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator2(text, "dist", rankMathFunction) },

		func(text string) *parserNode { return tryParseOperator1(text, "gauss", rankTermFunction) },

		func(text string) *parserNode { return tryParseOperator1(text, "deriv", rankTermFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "inti", rankTermFunction) },

		func(text string) *parserNode { return tryParseTerm(text) },
		func(text string) *parserNode { return tryParseVaraible(text) },
	)

	customeChecks = append(customeChecks,
		termCheck,
	)
}

type parserNode struct {
	*node
	parserChilds []*parserNode
	rank         int
	minChilds    int // -1 -> infinite
	maxChilds    int
}

func newParserNode(rank int, maxChilds int, minChilds int, node *node) *parserNode {
	return &parserNode{
		node: node,

		rank:      rank,
		minChilds: maxChilds,
		maxChilds: minChilds,
	}
}

func tryParseOperator1(text string, name string, rank int) *parserNode {
	if text != name {
		return nil
	}

	node := newParserNode(rank, 1, 1, newNode(name, 0, flagAction, flagOperator1))
	return node
}
func tryParseOperator2(text string, name string, rank int) *parserNode {
	if text != name {
		return nil
	}

	node := newParserNode(rank, 2, 2, newNode(name, 0, flagAction, flagOperator2))
	return node
}
func tryParseNumber(text string) *parserNode {
	if isNumber(text) {
		if x, err := strconv.ParseFloat(text, 64); !handelError(err) {
			return newParserNode(rankTermEnd, 0, 0, newNode(text, x, flagData, flagNumber))
		}
	}
	return nil
}

var currentVariables []*node

func tryParseVaraible(text string) *parserNode {
	for _, variable := range currentVariables {
		if text == variable.data {
			return newParserNode(rankTermEnd, 0, 0, newNode(text, 0, flagData, flagVariable))
		}
	}
	return nil
}

func tryParseTerm(text string) *parserNode {
	for _, term := range terms {
		if text == term.name {
			return newParserNode(rankTerm, len(term.variables), len(term.variables), newNode(text, 0, flagAction, flagTerm))
		}
	}
	return nil
}

func tryParseReplaceRulePart(text string) *parserNode {
	parts := splitAny(text, "_")
	if len(parts) == 2 {
		switch parts[0] {
		case "all":
			return newParserNode(rankTermEnd, 0, 0, newNode(parts[1], 0, flagRuleData))
		case "data":
			return newParserNode(rankTermEnd, 0, 0, newNode(parts[1], 0, flagData, flagRuleData))
		case "num":
			return newParserNode(rankTermEnd, 0, 0, newNode(parts[1], 0, flagNumber, flagRuleData))
		case "var":
			return newParserNode(rankTermEnd, 0, 0, newNode(parts[1], 0, flagVariable, flagRuleData))
		case "const":
			return newParserNode(rankTermEnd, 0, 0, newNode(parts[1], 0, flagConstant, flagRuleData))

		}
	}
	return nil
}

func (p *parserNode) updateChilds() {
	p.childs = make([]*node, len(p.parserChilds))
	for i, child := range p.parserChilds {
		child.updateChilds()
		p.childs[i] = child.node
	}
}

var customeChecks []func(p *parserNode) error

func (p *parserNode) check() error {
	for _, child := range p.parserChilds {
		err := child.check()
		if err != nil {
			return err
		}
	}

	for _, check := range customeChecks {
		err := check(p)
		if err != nil {
			return err
		}
	}

	if len(p.parserChilds) < p.minChilds {
		return newError(errorTypParsing, errorCriticalLevelPartial, "Node: \""+p.data+"\" has not enought children!")
	}
	if len(p.parserChilds) > p.maxChilds {
		return newError(errorTypParsing, errorCriticalLevelPartial, "Node: \""+p.data+"\" has to many children!")
	}
	return nil
}

func parseRoot(parseFuncs []func(text string) *parserNode, parts ...string) (*parserNode, int, error) {
	rootVar := newParserNode(0, 1, 1, newNode("", 0))
	root := &rootVar
	currentVar := newParserNode(0, 1, 1, newNode("", 0))
	current := &currentVar
	*current = rootVar

	var i int
	for i = 0; i < len(parts); i++ {
		part := parts[i]

		// When hitting a bracket we go one recursion deeper and parse there until we hit a closing bracket.
		// And skip all parts that have been parsed in the bracket.
		if part == "(" {
			subRoot, index, err := parseRoot(parseFuncs, parts[i+1:]...)
			if err != nil {
				return *root, i, err
			}

			(*subRoot).setFlag(flagBracketRoot, true)
			(*current).parserChilds = append((*current).parserChilds, subRoot)

			i += index + 1
			continue
		}
		if part == ")" {
			break
		}

		// Trying to parse the string
		node, err := tryParse(part, parseFuncs)
		if err != nil {
			return *root, i, err
		}

		addParsedNode(node, root, current)
	}

	err := (*root).check()
	if err != nil {
		return *root, i, err
	}

	(*root).updateChilds()

	return *root, i, nil
}
func tryParse(part string, parseFuncs []func(text string) *parserNode) (node *parserNode, err error) {
	for _, parseFunc := range parseFuncs {
		node = parseFunc(part)
		if node != nil {
			return node, nil
		}
	}
	return nil, newError(errorTypParsing, errorCriticalLevelPartial, "Expression: \""+part+"\" could not be parsed.")
}

func (p *parserNode) setParserChilds(childs ...*parserNode) {
	p.parserChilds = childs
}

// addParsedNode adds the node to the tree, the rank of the node is used to detemine where the node is placed.
func addParsedNode(newNode *parserNode, root **parserNode, current **parserNode) {

	// Case one the current newNode is noen so just replace it.
	if (*current).hasFlag(flagNone) {
		newNode.setParserChilds((*current).parserChilds...)
		partent := getParentOfNode(*current, *root)
		if partent != nil {
			for i, child := range partent.parserChilds {
				if child == (*current) {
					partent.parserChilds[i] = newNode
				}
			}
		}
		*root = newNode
		*current = newNode

		// Case two the newNode rank is higer the current so the newNode needs to be child of current.
	} else if newNode.rank > (*current).rank || newNode.hasFlag(flagBracketRoot) {

		childs := (*current).parserChilds
		if len(childs) < (*current).maxChilds {
			childs = append(childs, newNode)

		} else {
			// current is full so we need to push down one child so there is sapce for the newNode
			child := childs[len(childs)-1]
			childs[len(childs)-1] = newNode

			newNode.setParserChilds(child)
		}
		(*current).setParserChilds(childs...)

		// set current as newNode when we adden an action newNode
		if newNode.hasFlag(flagAction) {
			*current = newNode
		}

	} else if newNode.rank <= (*current).rank {

		// There is an empti parent
		partent := getParentOfNode(*current, *root)
		if partent == nil {
			newNode.setParserChilds(*current)
			*root = newNode
			*current = newNode
		} else {

			// We need to look higher in the tree.
			*current = partent
			addParsedNode(newNode, root, current)

		}
	}
}
func getParentOfNode(ofNode *parserNode, root *parserNode) *parserNode {

	var recursivSearch func(*parserNode, *parserNode) (bool, *parserNode)
	recursivSearch = func(current *parserNode, seachNode *parserNode) (foundNode bool, parent *parserNode) {
		if current == seachNode {
			return true, nil
		}

		for _, child := range current.parserChilds {
			foundNode, parent = recursivSearch(child, seachNode)
			if foundNode {
				if parent == nil {
					return true, current
				}
				return true, parent
			}
		}

		return false, nil
	}

	_, partent := recursivSearch(root, ofNode)
	return partent
}

package V4

const (
	rankNone            = 0
	rankRoot            = 1
	rankAppend          = 2
	rankAddSub          = 3
	rankMul             = 4
	rankPow             = 5
	rankMathFunction    = 6
	rankSubOpperation   = 7
	rankTerm            = 8
	rankComplexFunction = 9
	rankTermEnd         = 20
)

type parserNode struct {
	*node
	parserChilds []*parserNode
	rank         int
	minChilds    int // -1 -> infinite
	maxChilds    int
}

func NewParserNode(rank int, maxChilds int, minChilds int, node *node) *parserNode {
	return &parserNode{
		node: node,

		rank:      rank,
		minChilds: maxChilds,
		maxChilds: minChilds,
	}
}
func (p *parserNode) setParserChilds(childs ...*parserNode) {
	p.parserChilds = childs
}
func (p *parserNode) updateChilds() {
	p.childs = make([]*node, len(p.parserChilds))
	for i, child := range p.parserChilds {
		child.updateChilds()
		p.childs[i] = child.node
	}
}

var tryParseFuncs = []func(text string) *parserNode{
	func(text string) *parserNode { return tryParseNumber(text) },
	func(text string) *parserNode { return tryParseVaraible(text) },
	func(text string) *parserNode { return tryParseSimpleOpperator(text, "+", rankAddSub) },
	func(text string) *parserNode { return tryParseSimpleOpperator(text, "-", rankAddSub) },
	func(text string) *parserNode { return tryParseSimpleOpperator(text, "*", rankMul) },
	func(text string) *parserNode { return tryParseSimpleOpperator(text, "/", rankMul) },
	func(text string) *parserNode { return tryParseSimpleOpperator(text, "pow", rankPow) },
}

func tryParseSimpleOpperator(text string, name string, rank int) *parserNode {
	if text != name {
		return nil
	}

	node := NewParserNode(rank, 2, 2, NewNode(name, nodeFlagAction, nodeFlagOpperator))
	return node
}
func tryParseNumber(text string) *parserNode {
	if !isNumber(text) {
		return nil
	}

	node := NewParserNode(rankTermEnd, 0, 0, NewNode(text, nodeFlagData, nodeFlagNumber))
	return node
}

var currentVariables []*node

func tryParseVaraible(text string) *parserNode {
	for _, variable := range currentVariables {
		if text == variable.data {
			return NewParserNode(rankTermEnd, 0, 0, NewNode(text, nodeFlagData, nodeFlagVariable))
		}
	}
	return nil
}

func parseRoot(parts ...string) (*parserNode, int, error) {
	rootVar := NewParserNode(0, 1, 1, NewNode(""))
	root := &rootVar
	currentVar := NewParserNode(0, 1, 1, NewNode(""))
	current := &currentVar

	var i int
	for i = 0; i < len(parts); i++ {
		part := parts[i]

		// When hitting a bracket we go one recursion deeper and parse there until we hit a closing bracket.
		// And skip all parts that have been parsed in the bracket.
		if part == "(" {
			subRoot, index, err := parseRoot(parts[i+1:]...)
			if err != nil {
				return *root, i, err
			}

			(*subRoot).setFlag(nodeFlagBracketRoot, true)
			(*current).parserChilds = append((*current).parserChilds, subRoot)
			i += index + 1
			continue
		}
		if part == ")" {
			break
		}

		// Trying to parse the string
		node, err := tryParse(part)
		if err != nil {
			return *root, i, err
		}

		addParsedNode(node, root, current)
	}

	(*root).updateChilds()

	return *root, i, nil
}
func tryParse(part string) (node *parserNode, err error) {
	for _, parseFunc := range tryParseFuncs {
		node = parseFunc(part)
		if node != nil {
			return node, nil
		}
	}
	return nil, newError(errorTypParsing, errorCriticalLevelPartial, "Expression: \""+part+"\" could not be parsed.")
}

// addParsedNode adds the node to the tree, the rank of the node is used to detemine where the node is placed.
func addParsedNode(newNode *parserNode, root **parserNode, current **parserNode) {

	// Case one the current newNode is noen so just replace it.
	if (*current).hasFlag(nodeFlagNone) {
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
	} else if newNode.rank > (*current).rank || newNode.hasFlag(nodeFlagBracketRoot) {

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
		if newNode.hasFlag(nodeFlagAction) {
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

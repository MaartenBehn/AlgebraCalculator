package V4

const (
	flagNone = 0
	flagRoot = 1

	// Basic Types no Children
	flagData     = 10
	flagNumber   = 11
	flagConstant = 12
	flagVariable = 13

	// Action Typs have Children
	flagAction    = 20
	flagOperator1 = 21
	flagOperator2 = 22
	flagFraction  = 23

	flagBracketRoot = 30

	flagMax = 40
)

type node struct {
	childs []*node

	data       string
	dataNumber float64

	flagValues [flagMax]bool
}

func NewNode(data string, flags ...int) *node {
	node := &node{
		data: data,
	}

	// Set all flagValues that are parsed in.
	for _, flag := range flags {
		node.setFlag(flag, true)
	}

	// If no flagValues are set node will ge flag none
	if len(flags) == 0 {
		node.setFlag(flagNone, true)
	}

	return node
}
func (n *node) setChilds(childs ...*node) {
	n.childs = childs
}
func (n *node) setFlag(flag int, value bool) {
	// Remove none when an other flag is set.
	if n.flagValues[flagNone] && flag != flagNone && value {
		n.flagValues[flagNone] = false
	}

	n.flagValues[flag] = value
}
func (n *node) hasFlag(flag int) bool {
	return n.flagValues[flag]
}
func (n *node) hasAllFlagsOfNode(reference *node) bool {
	for flag, flagValue := range reference.flagValues {
		if flagValue && !n.hasFlag(flag) {
			return false
		}
	}
	return true
}
func (n *node) hasAllFlagsOfNodeDeep(reference *node) bool {
	if !n.hasAllFlagsOfNode(reference) {
		return false
	}

	for i, child := range n.childs {
		if i >= len(reference.childs) {
			return false
		}
		if !child.hasAllFlagsOfNodeDeep(reference.childs[i]) {
			return false
		}
	}
	return true
}

package V4

const (
	nodeFlagNone = 0
	nodeFlagRoot = 1

	// Basic Types no Children
	nodeFlagData     = 10
	nodeFlagNumber   = 11
	nodeFlagConstant = 12
	nodeFlagVariable = 13

	// Action Typs have Children
	nodeFlagAction       = 20
	nodeFlagOpperator    = 21
	nodeFlagMathFunction = 22

	nodeFlagBracketRoot = 30

	nodeFlagMax = 100
)

type node struct {
	childs []*node

	data  string
	flags [nodeFlagMax]bool
}

func NewNode(data string, flags ...int) *node {
	node := &node{
		data: data,
	}

	// Set all flags that are parsed in.
	for _, flag := range flags {
		node.setFlag(flag, true)
	}

	// If no flags are set node will ge flag none
	if len(flags) == 0 {
		node.setFlag(nodeFlagNone, true)
	}

	return node
}

func (n *node) setChilds(childs ...*node) {
	n.childs = childs
}
func (n *node) setFlag(flag int, value bool) {
	// Remove none when an other flag is set.
	if n.flags[nodeFlagNone] && flag != nodeFlagNone && value {
		n.flags[nodeFlagNone] = false
	}

	n.flags[flag] = value
}
func (n *node) hasFlag(flag int) bool {
	return n.flags[flag]
}

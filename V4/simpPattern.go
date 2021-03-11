package V4

import "math"

type simpPattern struct {
	pattern func(root *node) bool
	apply   func(root *node) *node
}

var simpPatterns = []simpPattern{
	solveOperator2("+", func(x float64, y float64) float64 { return x + y }),
	solveOperator2Edge("+", func(x float64, y float64) float64 { return x + y }),

	solveOperator2("-", func(x float64, y float64) float64 { return x - y }),
	solveOperator2Edge("-", func(x float64, y float64) float64 { return x - y }),

	solveOperator2("*", func(x float64, y float64) float64 { return x * y }),
	solveOperator2Edge("*", func(x float64, y float64) float64 { return x * y }),

	solveOperator2("/", func(x float64, y float64) float64 { return x / y }),
	solveOperator2Edge("/", func(x float64, y float64) float64 { return x / y }),

	solveOperator1("sin", math.Sin),
	solveOperator1("sinh", math.Sinh),
	solveOperator1("asin", math.Asin),
	solveOperator1("asinh", math.Asinh),

	solveOperator1("cos", math.Cos),
	solveOperator1("cosh", math.Cosh),
	solveOperator1("acos", math.Acos),
	solveOperator1("acosh", math.Acosh),

	solveOperator1("tan", math.Tan),
	solveOperator1("tanh", math.Tanh),
	solveOperator1("atan", math.Atan),
	solveOperator1("atanh", math.Atanh),
	solveOperator2("atan2", math.Atan2),
}

func solveOperator1(name string, function func(x float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(flagOperator1) &&
				root.childs[0].hasFlag(flagNumber)
		},
		func(root *node) *node {
			node := NewNode("", flagData, flagNumber)
			node.dataNumber = function(root.childs[0].dataNumber)
			return node
		},
	}
}
func solveOperator2(name string, function func(x float64, y float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(flagOperator2) &&
				root.childs[0].hasFlag(flagNumber) &&
				root.childs[1].hasFlag(flagNumber)
		},
		func(root *node) *node {
			node := NewNode("", flagData, flagNumber)
			node.dataNumber = function(root.childs[0].dataNumber, root.childs[1].dataNumber)
			return node
		},
	}
}
func solveOperator2Edge(name string, function func(x float64, y float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(flagOperator2) &&
				root.childs[0].data == name &&
				root.childs[0].hasFlag(flagOperator2) &&
				root.childs[1].hasFlag(flagNumber) &&
				root.childs[0].childs[1].hasFlag(flagNumber)
		},
		func(root *node) *node {
			node := NewNode("", flagData, flagNumber)
			node.dataNumber = function(root.childs[0].childs[1].dataNumber, root.childs[1].dataNumber)
			edge := NewNode(name, flagAction, flagOperator2)
			edge.setChilds(root.childs[0].childs[0], node)
			return edge
		},
	}
}

func simpReplace(searchRoot *node, replaceRoot *node) simpPattern {

	var match func(*node, *node) bool
	match = func(current *node, search *node) bool {
		if !current.hasAllFlagsOfNode(search) {
			return false
		}

		if search.hasFlag(flagAction) && current.data != search.data {
			return false
		}

		for i, child := range current.childs {
			if i >= len(search.childs) {
				return false
			}
			if !match(child, search.childs[i]) {
				return false
			}
		}
		return true
	}

	var dataNodes []*node
	var findDataNodes func(*node, *node)
	findDataNodes = func(current *node, search *node) {
		for i, child := range current.childs {
			findDataNodes(child, search.childs[i])
		}
		if len(search.data) == 1 && search.data[0] == '$' {
			id := int(search.dataNumber)

			childs := current.childs
			*current = *dataNodes[id]
			current.childs = childs
		}
	}

	var replace func(*node, *node)
	replace = func(current *node, replacement *node) {
		for i, child := range current.childs {
			replace(child, replacement.childs[i])
		}
		if len(replacement.data) == 1 && replacement.data[0] == '$' {
			id := int(replacement.dataNumber)

			childs := current.childs
			*current = *dataNodes[id]
			current.childs = childs

		} else {
			*current = *replacement
		}
	}

	return simpPattern{
		func(root *node) bool {
			return match(root, searchRoot)
		},
		func(root *node) *node {
			findDataNodes(root, searchRoot)
			replace(root, replaceRoot)
			return root
		},
	}
}

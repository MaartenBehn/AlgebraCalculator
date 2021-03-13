package V4

import "math"

func initSolve() {
	simpPatterns = append(simpPatterns,
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
	)
}

func solveOperator1(name string, function func(x float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(flagOperator1) &&
				root.childs[0].hasFlag(flagNumber)
		},
		func(root *node) *node {
			node := newNode("", function(root.childs[0].dataNumber), flagData, flagNumber)
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
			node := newNode("", function(root.childs[0].dataNumber, root.childs[1].dataNumber), flagData, flagNumber)
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
			node := newNode("", 0, flagData, flagNumber)
			node.dataNumber = function(root.childs[0].childs[1].dataNumber, root.childs[1].dataNumber)
			edge := newNode(name, 0, flagAction, flagOperator2)
			edge.setChilds(root.childs[0].childs[0], node)
			return edge
		},
	}
}

package V4

import "math"

type simpPattern struct {
	pattern func(root *node) bool
	apply   func(root *node) *node
}

var simpPatterns = []simpPattern{
	simpleOperatorPattern("+", func(x float64, y float64) float64 { return x + y }),
	simpleOperatorPattern("-", func(x float64, y float64) float64 { return x - y }),
	simpleOperatorPattern("*", func(x float64, y float64) float64 { return x * y }),
	simpleOperatorPattern("/", func(x float64, y float64) float64 { return x / y }),

	simpleMathFunctionPattern("sin", math.Sin),
	simpleMathFunctionPattern("sinh", math.Sinh),
	simpleMathFunctionPattern("asin", math.Asin),
	simpleMathFunctionPattern("asinh", math.Asinh),

	simpleMathFunctionPattern("cos", math.Cos),
	simpleMathFunctionPattern("cosh", math.Cosh),
	simpleMathFunctionPattern("acos", math.Acos),
	simpleMathFunctionPattern("acosh", math.Acosh),

	simpleMathFunctionPattern("tan", math.Tan),
	simpleMathFunctionPattern("tanh", math.Tanh),
	simpleMathFunctionPattern("atan", math.Atan),
	simpleMathFunctionPattern("atanh", math.Atanh),
	simpleOperatorPattern("atan2", math.Atan2),
}

func simpleOperatorPattern(name string, function func(x float64, y float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(nodeFlagOpperator) &&
				root.childs[0].hasFlag(nodeFlagNumber) &&
				root.childs[1].hasFlag(nodeFlagNumber)
		},
		func(root *node) *node {
			node := NewNode("", nodeFlagData, nodeFlagNumber)
			node.dataFloat = function(root.childs[0].dataFloat, root.childs[1].dataFloat)
			return node
		},
	}
}

func simpleMathFunctionPattern(name string, function func(x float64) float64) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.data == name &&
				root.hasFlag(nodeFlagMathFunction) &&
				root.childs[0].hasFlag(nodeFlagNumber)
		},
		func(root *node) *node {
			node := NewNode("", nodeFlagData, nodeFlagNumber)
			node.dataFloat = function(root.childs[0].dataFloat)
			return node
		},
	}
}

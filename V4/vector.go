package V4

import (
	"math"
)

func initVector() {
	simpPatterns = append(simpPatterns,
		simpPattern{
			func(root *node) bool {
				return root.hasFlag(flagOperator2) && root.data == ","
			},
			func(root *node) *node {
				return vectorOpperatorToNode(root)
			},
			"\",\" operator to vector",
		},

		vectorMergeOperator2("+"),
		vectorMergeOperator2("-"),

		vectorMergeOperator1("sin"), // TODO check if you can actually do that.

		vectorApplyScalar("*"),

		vectorOperator2("dot", dot),
		vectorOperator1("len", magnitude),
		vectorOperator2("dist", dist),

		vectorSubOperation(),
	)
}

func newVector() *node {
	return newNode("Vector", 0, flagAction, flagVector)
}

func vectorOpperatorToNode(node *node) *node {
	vector := newVector()

	for _, child := range node.childs {
		if child.data == "," {
			vector.childs = append(vector.childs, vectorOpperatorToNode(child).childs...)
		} else {
			vector.childs = append(vector.childs, child)
		}
	}
	return vector
}

func vectorMergeOperator1(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator1) && root.data == name &&
				root.childs[0].hasFlag(flagVector)
		},
		func(root *node) *node {
			result := newVector()

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)
			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator1)
				result.childs[i].setChilds(root.childs[0].childs[i])
			}
			return result
		},
		"Vector merge " + name,
	}
}
func vectorMergeOperator2(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				root.childs[0].hasFlag(flagVector) &&
				root.childs[1].hasFlag(flagVector) &&
				len(root.childs[0].childs) == len(root.childs[1].childs)
		},
		func(root *node) *node {
			result := newVector()

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)
			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator2)
				result.childs[i].setChilds(root.childs[0].childs[i], root.childs[1].childs[i])
			}
			return result
		},
		"Vector merge " + name,
	}
}

func vectorApplyScalar(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				(root.childs[0].hasFlag(flagVector) && root.childs[1].hasFlag(flagData))
		},
		func(root *node) *node {
			result := newVector()

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)

			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator2)
				result.childs[i].setChilds(root.childs[0].childs[i], root.childs[1].copyDeep())
			}
			return result
		},
		"Apply scala " + name,
	}
}

func vectorOperator1(name string, function func(x *node) *node) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator1) && root.data == name &&
				root.childs[0].hasFlag(flagVector)
		},
		func(root *node) *node {
			return function(root.childs[0])
		},
		"Vector solve " + name,
	}
}
func vectorOperator2(name string, function func(x *node, y *node) *node) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				root.childs[0].hasFlag(flagVector) &&
				root.childs[1].hasFlag(flagVector) &&
				len(root.childs[0].childs) == len(root.childs[1].childs)
		},
		func(root *node) *node {
			return function(root.childs[0], root.childs[1])
		},
		"Vector solve " + name,
	}
}

func dot(x *node, y *node) *node {
	result := newNode("+", 0, flagAction, flagOperator2)
	current := &result

	for i := 2; i < len(x.childs); i++ {
		(*current).setChilds(newNode("+", 0, flagAction, flagOperator2))
		current = &((*current).childs[0])
	}

	current = &result
	for i := len(x.childs) - 1; i >= 0; i-- {
		mul := newNode("*", 0, flagAction, flagOperator2)
		mul.setChilds(x.childs[i], y.childs[i])

		(*current).childs = append((*current).childs, mul)
		if i > 1 {
			current = &((*current).childs[0])
		}
	}

	return result
}
func magnitude(x *node) *node {
	result := newNode("sqrt", 0, flagAction, flagOperator1)
	current := &result

	for i := 1; i < len(x.childs); i++ {
		(*current).setChilds(newNode("+", 0, flagAction, flagOperator2))
		current = &((*current).childs[0])
	}

	current = &result.childs[0]
	for i := len(x.childs) - 1; i >= 0; i-- {
		mul := newNode("pow", 0, flagAction, flagOperator2)
		mul.setChilds(x.childs[i], newNode("", 2, flagData, flagNumber))

		(*current).childs = append((*current).childs, mul)
		if i > 1 {
			current = &((*current).childs[0])
		}
	}

	return result
}
func dist(x *node, y *node) *node {
	input := newVector()

	input.childs = make([]*node, len(x.childs))
	for i := range x.childs {
		input.childs[i] = newNode("-", 0, flagAction, flagOperator2)
		input.childs[i].setChilds(y.childs[i], x.childs[i])
	}

	result := newNode("abs", 0, flagAction, flagOperator1)
	result.setChilds(magnitude(input))

	return result
}

func vectorSubOperation() simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == "." &&
				root.childs[0].hasFlag(flagVector) &&
				root.childs[1].hasFlag(flagNumber) &&
				root.childs[1].dataNumber == math.Trunc(root.childs[1].dataNumber)
		},
		func(root *node) *node {
			result := newVector()

			number := int(root.childs[1].dataNumber)

			for number > 0 {
				digit := number % 10
				number /= 10

				if digit > len(root.childs[0].childs) || digit <= 0 {
					handelError(newError(errorTypParsing, errorCriticalLevelNon, "SubOperation index out of bounds."))
					return root.childs[0]
				}

				result.childs = append([]*node{root.childs[0].childs[digit-1]}, result.childs...)
			}

			return result
		},
		"Solve SubOperation",
	}
}

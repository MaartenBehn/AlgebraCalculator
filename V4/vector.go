package V4

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
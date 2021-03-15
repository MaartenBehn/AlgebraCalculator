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
	)
}

func vectorOpperatorToNode(node *node) *node {
	vector := newNode("Vector", 0, flagAction, flagVector)

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
			result := newNode("Vector", 0, flagAction, flagVector)

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)
			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator1)
				result.childs[i].setChilds(root.childs[0].childs[i])
			}
			return result
		},
		"Vector solve " + name,
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
			result := newNode("Vector", 0, flagAction, flagVector)

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)
			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator2)
				result.childs[i].setChilds(root.childs[0].childs[i], root.childs[1].childs[i])
			}
			return result
		},
		"Vector solve " + name,
	}
}

func vectorApplyScalar(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				(root.childs[0].hasFlag(flagVector) && root.childs[1].hasFlag(flagData))
		},
		func(root *node) *node {
			result := newNode("Vector", 0, flagAction, flagVector)

			dimensions := len(root.childs[0].childs)
			result.childs = make([]*node, dimensions)

			for i := 0; i < dimensions; i++ {
				result.childs[i] = newNode(name, 0, flagAction, flagOperator2)
				result.childs[i].setChilds(root.childs[0].childs[i], root.childs[1].copyDeep())
			}
			return result
		},
		"Vector solve " + name,
	}
}

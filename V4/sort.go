package V4

func initSort() {
	simpPatterns = append(simpPatterns,
		sortOp2("+"),
		sortOp2Edge("+"),
		sortOp2("*"),
		sortOp2Edge("*"),
	)
}

func sortOp2(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				root.childs[0].data != name && root.childs[1].data != name &&
				root.childs[0].getIdentityDeep() > root.childs[1].getIdentityDeep()
		},
		func(root *node) *node {
			root.childs = []*node{root.childs[1], root.childs[0]}
			return root
		},
		"Sort " + name,
	}

}

func sortOp2Edge(name string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator2) && root.data == name &&
				root.childs[0].hasFlag(flagOperator2) && root.childs[0].data == name &&
				root.childs[0].childs[1].data != name && root.childs[1].data != name &&
				root.childs[0].childs[1].getIdentityDeep() > root.childs[1].getIdentityDeep()
		},
		func(root *node) *node {
			child1 := root.childs[1]
			root.childs[1] = root.childs[0].childs[1]
			root.childs[0].childs[1] = child1
			return root
		},
		"Sort Edge " + name,
	}

}

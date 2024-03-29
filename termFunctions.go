package AlgebraCalculator

func initTermFunctions() {
	simpPatterns = append(simpPatterns,
		termFunctionGauss(),
	)
}

func termFindBaseChilds(node *node, childs *[]*node) {
	if !(node.hasFlag(flagOperator2) && node.data == "+") {
		return
	}

	for _, child := range node.childs {
		if child.hasFlag(flagOperator2) && child.data == "+" {
			termFindBaseChilds(child, childs)
		} else {
			*childs = append(*childs, child)
		}
	}
}

func termFunctionGauss() simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator1) && root.data == "gauss"
		},
		func(root *node) *node {
			dimentions := 1

			if root.childs[0].hasFlag(flagVector) {
				dimentions = len(root.childs[0].childs)
			}

			var roots []*node
			if dimentions != 1 {
				roots = root.childs[0].childs
			} else {
				roots = []*node{root.childs[0]}
			}

			var varsList []struct {
				name    string
				ammount []float64
			}
			var numsList []float64

			for _, root := range roots {
				// Finind childs of +ese
				var childs []*node
				termFindBaseChilds(root, &childs)

				var nums []float64

				for _, child := range childs {
					if child.hasFlag(flagOperator2) && child.data == "*" &&
						child.childs[0].hasFlag(flagNumber) &&
						child.childs[1].hasFlag(flagVariable) {

						found := false
						for i, s := range varsList {
							if s.name == child.childs[1].data {
								varsList[i].ammount = append(varsList[i].ammount, child.childs[0].dataNumber)
								found = true
							}
						}

						if !found {
							varsList = append(varsList, struct {
								name    string
								ammount []float64
							}{name: child.childs[1].data, ammount: []float64{child.childs[0].dataNumber}})
						}

					} else if child.hasFlag(flagNumber) {
						nums = append(nums, child.dataNumber)
					}
				}

				if len(nums) > 1 {
					handelError(newError(errorTypSolving, errorCriticalLevelNon, "More than one single Scalar found! "))
				} else if len(nums) == 0 {
					nums = append(nums, 0)
				}

				numsList = append(numsList, nums[0])
			}

			a := make([][]float64, dimentions)
			for i := 0; i < dimentions; i++ {
				a[i] = make([]float64, dimentions)
			}
			for i := 0; i < dimentions; i++ {
				for j := 0; j < dimentions; j++ {

					a[j][i] = varsList[i].ammount[j]
				}
			}

			x, err := gaussPartial(a, numsList)
			handelError(err)

			result := newVector()
			result.childs = make([]*node, len(x))
			for i, f := range x {
				result.childs[i] = newNode("", f*-1, flagData, flagNumber)
			}

			return result
		},
		"Gaussian elimination",
	}
}

func termDerivative() simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator1) && root.data == "deriv"
		},
		func(root *node) *node {

			var childs []*node
			termFindBaseChilds(root.childs[0], &childs)

			for _, child := range childs {
				if child.hasFlag(flagOperator2) && child.data == "*" && child.childs[0].hasFlag(flagNumber) {
					if child.childs[1].hasFlag(flagOperator2) && child.childs[1].data == "pow" {

						instertNode := newNode("*", 0, flagAction, flagOperator2)
						instertNode.childs = []*node{
							newNode("number", child.childs[1].childs[1].dataNumber, flagData, flagNumber),
							newNode("number", child.childs[0].dataNumber, flagData, flagNumber),
						}
						child.childs[0] = instertNode

						instertNode2 := newNode("+", 0, flagAction, flagOperator2)
						instertNode2.childs = []*node{
							child.childs[1].childs[1],
							newNode("number", 1, flagData, flagNumber),
						}
						child.childs[1].childs[1] = instertNode2

					} else {

						instertNode := newNode("pow", 0, flagAction, flagOperator2)
						instertNode.childs = []*node{
							child.childs[1],
							newNode("number", 2, flagData, flagNumber),
						}
						child.childs[1] = instertNode
					}
				}
			}

			return root
		},
		"Derivative",
	}
}

package AlgebraCalculator

func initTermFunctions() {
	simpPatterns = append(simpPatterns,
		termFunctionGauss(),
	)
}

func termFunctionGauss() simpPattern {
	return simpPattern{
		func(root *node) bool {
			return root.hasFlag(flagOperator1) && root.data == "gauss" &&
				root.childs[0].hasFlag(flagTerm)
		},
		func(root *node) *node {
			var term *term
			for _, t := range terms {
				if t.name == root.childs[0].data {
					term = t
					break
				}
			}

			dimentions := len(term.variables)

			var variableVectors []*node
			gaussFindVaribleVectors(term.root, &variableVectors)
			if len(variableVectors) != dimentions {
				handelError(newError(errorTypSolving, errorCriticalLevelNon, "Not all or more variables found than in term."))
				return nil
			}

			a := make([][]float64, dimentions)
			for i := 0; i < dimentions; i++ {
				a[i] = make([]float64, len(variableVectors))
			}
			for i := 0; i < dimentions; i++ {
				for j := 0; j < dimentions; j++ {
					a[j][i] = variableVectors[i].childs[0].childs[j].dataNumber
				}
			}

			var vectors []*node
			gaussFindVectors(term.root, &vectors)
			if len(vectors) != 1 {
				handelError(newError(errorTypSolving, errorCriticalLevelNon, "Not exactly one b vector found."))
				return nil
			}
			b := make([]float64, dimentions)
			for i := 0; i < dimentions; i++ {
				b[i] = vectors[0].childs[0].childs[i].dataNumber
			}

			x, err := gaussPartial(a, b)
			handelError(err)

			result := newVector()

			result.childs = make([]*node, len(x))
			for i, termVariable := range term.variables {
				for j, variableVector := range variableVectors {
					if termVariable.data == variableVector.childs[1].data {
						result.childs[i] = newNode("", x[j]*-1, flagData, flagNumber)
					}
				}
			}

			return result
		},
		"Gaussian elimination",
	}
}

func gaussFindVaribleVectors(node *node, vectors *[]*node) {
	for _, child := range node.childs {
		gaussFindVaribleVectors(child, vectors)
	}

	if node.hasFlag(flagOperator2) &&
		node.data == "*" &&
		node.childs[0].hasFlag(flagVector) &&
		node.childs[1].hasFlag(flagVariable) {
		*vectors = append(*vectors, node)
	}
}
func gaussFindVectors(node *node, vectors *[]*node) {
	for _, child := range node.childs {
		gaussFindVectors(child, vectors)
	}

	if node.hasFlag(flagOperator2) &&
		node.data == "+" &&
		node.childs[0].hasFlag(flagVector) {
		*vectors = append(*vectors, node)
	}
}

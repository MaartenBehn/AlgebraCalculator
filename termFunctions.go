package AlgebraCalculator

import (
	"AlgebraCalculator/log"
)

var termFunctions = []*termFunction{
	newTermFunction("gauss", gauss),
	newTermFunction("removeSub", removeSubOpperations),
}

type termFunction struct {
	*namedNode
	function func(term *term) (iNode, error)
}

func newTermFunction(name string, function func(term *term) (iNode, error)) *termFunction {
	return &termFunction{
		namedNode: newNamedNode(newNode(typComplexFunction, rankComplexFunction, 1), name),
		function:  function,
	}
}

func (f *termFunction) copy() iNode {
	copy := newTermFunction(f.name, f.function)
	copy.childs = make([]iNode, len(f.childs))

	for i, child := range f.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (f *termFunction) solve() bool {

	if len(f.childs) != 1 {
		handelError(newError(errorTypSolving, errorCriticalLevelNon, "term function has 0 or more than 1 Argument."))
	}

	if f.childs[0].getType() != typTerm {
		handelError(newError(errorTypSolving, errorCriticalLevelNon, "Argumens is not a term."))
	}

	node, err := f.function(f.childs[0].(*term))
	handelError(err)
	if node != nil {
		insertNode(f, node)
		return true
	}
	return false
}
func (f *termFunction) print() {
	log.Print(f.name)
	if len(f.childs) > 0 {

		log.Print("<")
		for i, child := range f.childs {
			child.print()
			if i < len(f.childs)-1 {
				log.Print(" ")
			}
		}
		log.Print(">")
	}
}

func gauss(term *term) (iNode, error) {
	dimentions := len(term.variables)

	var variableVectors []iNode
	gaussFindVaribleVectors(term.root, &variableVectors)
	if len(variableVectors) != dimentions {
		return nil, newError(errorTypSolving, errorCriticalLevelNon, "Not all or more variables found than in term.")
	}

	a := make([][]float64, dimentions)
	for i := 0; i < dimentions; i++ {
		a[i] = make([]float64, len(variableVectors))
	}
	for i := 0; i < dimentions; i++ {
		for j := 0; j < dimentions; j++ {
			a[j][i] = variableVectors[i].getChilds()[0].(*vector).values[j]
		}
	}

	var vectors []iNode
	gaussFindVectors(term.root, &vectors)
	if len(vectors) != 1 {
		return nil, newError(errorTypSolving, errorCriticalLevelNon, "Not exactly one b vector found.")
	}
	b := make([]float64, dimentions)
	for i := 0; i < dimentions; i++ {
		b[i] = vectors[0].getChilds()[0].(*vector).values[i]
	}

	x, err := gaussPartial(a, b)
	handelError(err)

	result := make([]float64, len(x))
	for i, termVariable := range term.variables {
		for j, variableVector := range variableVectors {
			if termVariable.name == variableVector.getChilds()[1].(*variable).name {
				result[i] = x[j] * -1
			}
		}
	}

	return newVector(result), nil
}
func gaussFindVaribleVectors(node iNode, vectors *[]iNode) {
	for _, child := range node.getChilds() {
		gaussFindVaribleVectors(child, vectors)
	}

	if node.getType() == typOpperator &&
		node.(iNamedNode).getName() == "*" &&
		node.getChilds()[0].getType() == typVector &&
		node.getChilds()[1].getType() == typVariable {
		*vectors = append(*vectors, node)
	}
}
func gaussFindVectors(node iNode, vectors *[]iNode) {
	for _, child := range node.getChilds() {
		gaussFindVectors(child, vectors)
	}

	if node.getType() == typOpperator &&
		node.(iNamedNode).getName() == "+" &&
		node.getChilds()[0].getType() == typVector {
		*vectors = append(*vectors, node)
	}
}

func removeSubOpperations(term *term) (iNode, error) {
	removeSubOpperationsReplace(term.root)
	return term.root, nil
}
func removeSubOpperationsReplace(node iNode) {
	for _, child := range node.getChilds() {
		removeSubOpperationsReplace(child)
	}

	if node.getType() == typSubOperation {
		insertNode(node, node.getChilds()[0])
	}
}

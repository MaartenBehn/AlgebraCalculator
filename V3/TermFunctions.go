package V3

import (
	"fmt"
)

var termFunctions = []*TermFunction{
	NewTermFunction("gauss", gauss),
	NewTermFunction("removeSub", removeSubOpperations),
}

type TermFunction struct {
	*NamedNode
	function func(term *Term) (INode, error)
}

func NewTermFunction(name string, function func(term *Term) (INode, error)) *TermFunction {
	return &TermFunction{
		NamedNode: NewNamedNode(NewNode(TypComplexFunction, RankComplexFunction, 1), name),
		function:  function,
	}
}

func (f *TermFunction) copy() INode {
	copy := NewTermFunction(f.name, f.function)
	copy.childs = make([]INode, len(f.childs))

	for i, child := range f.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (f *TermFunction) solve() bool {

	if len(f.childs) != 1 {
		handelError(NewError(ErrorTypSolving, ErrorCriticalLevelNon, "Term function has 0 or more than 1 Argument."))
	}

	if f.childs[0].getType() != TypTerm {
		handelError(NewError(ErrorTypSolving, ErrorCriticalLevelNon, "Argumens is not a term."))
	}

	node, err := f.function(f.childs[0].(*Term))
	handelError(err)
	if node != nil {
		insertNode(f, node)
		return true
	}
	return false
}
func (f *TermFunction) print() {
	fmt.Print(f.name)
	if len(f.childs) > 0 {

		fmt.Print("<")
		for i, child := range f.childs {
			child.print()
			if i < len(f.childs)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print(">")
	}
}

func gauss(term *Term) (INode, error) {
	dimentions := len(term.variables)

	var variableVectors []INode
	gaussFindVaribleVectors(term.root, &variableVectors)
	if len(variableVectors) != dimentions {
		return nil, NewError(ErrorTypSolving, ErrorCriticalLevelNon, "Not all or more variables found than in term.")
	}

	a := make([][]float64, dimentions)
	for i := 0; i < dimentions; i++ {
		a[i] = make([]float64, len(variableVectors))
	}
	for i := 0; i < dimentions; i++ {
		for j := 0; j < dimentions; j++ {
			a[j][i] = variableVectors[i].getChilds()[0].(*Vector).values[j]
		}
	}

	var vectors []INode
	gaussFindVectors(term.root, &vectors)
	if len(vectors) != 1 {
		return nil, NewError(ErrorTypSolving, ErrorCriticalLevelNon, "Not exactly one b Vector found.")
	}
	b := make([]float64, dimentions)
	for i := 0; i < dimentions; i++ {
		b[i] = vectors[0].getChilds()[0].(*Vector).values[i]
	}

	x, err := GaussPartial(a, b)
	handelError(err)

	result := make([]float64, len(x))
	for i, variable := range term.variables {
		for j, variableVector := range variableVectors {
			if variable.name == variableVector.getChilds()[1].(*Variable).name {
				result[i] = x[j] * -1
			}
		}
	}

	return NewVector(result), nil
}
func gaussFindVaribleVectors(node INode, vectors *[]INode) {
	for _, child := range node.getChilds() {
		gaussFindVaribleVectors(child, vectors)
	}

	if node.getType() == TypOpperator &&
		node.(INamedNode).getName() == "*" &&
		node.getChilds()[0].getType() == TypVector &&
		node.getChilds()[1].getType() == TypVariable {
		*vectors = append(*vectors, node)
	}
}
func gaussFindVectors(node INode, vectors *[]INode) {
	for _, child := range node.getChilds() {
		gaussFindVectors(child, vectors)
	}

	if node.getType() == TypOpperator &&
		node.(INamedNode).getName() == "+" &&
		node.getChilds()[0].getType() == TypVector {
		*vectors = append(*vectors, node)
	}
}

func removeSubOpperations(term *Term) (INode, error) {
	removeSubOpperationsReplace(term.root)
	return term.root, nil
}
func removeSubOpperationsReplace(node INode) {
	for _, child := range node.getChilds() {
		removeSubOpperationsReplace(child)
	}

	if node.getType() == TypSubOperation {
		insertNode(node, node.getChilds()[0])
	}
}

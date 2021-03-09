package V3

import (
	"AlgebraCalculator/log"
	"math"
)

var mathFunctions = []*mathFunction{
	newMathFunction("sqrt", sqrt, 1),
	newMathFunction("degree", degree, 1),
	newMathFunction("radians", radians, 1),
	newMathFunction("sin", sin, 1),
	newMathFunction("cos", cos, 1),
	newMathFunction("tan", tan, 1),
	newMathFunction("len", magnitude, 1),
}

type mathFunction struct {
	*namedNode
	function func([]*vector) *vector
}

func newMathFunction(name string, function func([]*vector) *vector, attributeAmount int) *mathFunction {
	return &mathFunction{
		namedNode: newNamedNode(newNode(typMathFunction, rankMathFunction, attributeAmount), name),
		function:  function,
	}
}

func (f *mathFunction) copy() iNode {
	copy := newMathFunction(f.name, f.function, f.maxChilds)
	copy.childs = make([]iNode, len(f.childs))

	for i, child := range f.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (f *mathFunction) solve() bool {
	f.node.solve()

	var vectors []*vector
	for _, child := range f.childs {

		if child.getType() == typVector {
			vectors = append(vectors, child.(*vector))
		}
	}

	if len(vectors) == len(f.childs) {
		result := f.function(vectors)
		replaceNode(f, result)
		result.childs = nil
		return true
	}
	return false
}
func (f *mathFunction) print() {
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

func sqrt(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], math.Sqrt)
}

func degree(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (180.0 / math.Pi)
	})
}
func radians(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (math.Pi / 180.0)
	})
}

func sin(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], math.Sin)
}
func cos(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], math.Cos)
}
func tan(vectors []*vector) *vector {
	return genericOpperation1V(vectors[0], math.Tan)
}
func magnitude(vectors []*vector) *vector {
	var sum float64
	for _, value := range vectors[0].values {
		sum += math.Pow(value, 2)
	}
	sum = math.Sqrt(sum)
	return newVector([]float64{sum})
}

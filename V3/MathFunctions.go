package V3

import (
	"fmt"
	"math"
)

var mathFunctions = []*MathFunction{
	NewMathFunction("sqrt", sqrt, 1),
	NewMathFunction("degree", degree, 1),
	NewMathFunction("radians", radians, 1),
	NewMathFunction("sin", sin, 1),
	NewMathFunction("cos", cos, 1),
	NewMathFunction("tan", tan, 1),
}

type MathFunction struct {
	*NamedNode
	function func([]*Vector) *Vector
}

func NewMathFunction(name string, function func([]*Vector) *Vector, attributeAmount int) *MathFunction {
	return &MathFunction{
		NamedNode: NewNamedNode(NewNode(TypFunction, RankFunc, attributeAmount), name),
		function:  function,
	}
}

func (f *MathFunction) copy() INode {
	copy := NewMathFunction(f.name, f.function, f.maxChilds)
	copy.childs = make([]INode, len(f.childs))

	for i, child := range f.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (f *MathFunction) solve() {
	f.Node.solve()

	var vectors []*Vector
	for _, child := range f.childs {

		if child.getType() == TypVector {
			vectors = append(vectors, child.(*Vector))
		}
	}

	if len(vectors) == len(f.childs) {
		result := f.function(vectors)
		replaceNode(f, result)
		result.childs = nil
	}
}
func (f *MathFunction) print() {
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

func sqrt(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], math.Sqrt)
}

func degree(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (180.0 / math.Pi)
	})
}
func radians(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (math.Pi / 180.0)
	})
}

func sin(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], math.Sin)
}
func cos(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], math.Cos)
}
func tan(vectors []*Vector) *Vector {
	return genericOpperation1V(vectors[0], math.Tan)
}

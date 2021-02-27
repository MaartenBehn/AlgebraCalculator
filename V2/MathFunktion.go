package V2

import (
	"fmt"
	"math"
)

var mathFunctions = []MathFunction{
	{"sqrt", sqrt, 1},
	{"degree", degree, 1},
	{"radians", radians, 1},
	{"sin", sin, 1},
	{"cos", cos, 1},
	{"tan", tan, 1},
}

type MathFunction struct {
	name            string
	function        func([]Vector) Vector
	attributeAmount int
}

func (f MathFunction) getName() string {
	return f.name
}
func (f MathFunction) getType() int {
	return TypFunction
}
func (f MathFunction) getRank() int {
	return RankFunc
}
func (f MathFunction) isSolvable() bool {
	return true
}
func (f MathFunction) solve(term *Term, index int) bool {
	attributs := term.getSub(index+1, index+f.attributeAmount)

	areVectors := true
	for _, attribute := range attributs.parts {
		if attribute.getType() != TypVector {
			areVectors = false
		}
	}

	vectors := make([]Vector, len(attributs.parts))
	for i, attribute := range attributs.parts {
		vectors[i] = attribute.(Vector)
	}

	if areVectors {
		result := f.function(vectors)

		term.setSub(index, index+f.attributeAmount, NewTerm([]ITermPart{result}))
	}
	return false
}
func (f MathFunction) print() {
	fmt.Print(f.name)
}
func (f MathFunction) getSimplify() int {
	return SimplifyVariable
}

func sqrt(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], math.Sqrt)
}

func degree(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (180.0 / math.Pi)
	})
}
func radians(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], func(f float64) float64 {
		return f * (math.Pi / 180.0)
	})
}

func sin(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], math.Sin)
}
func cos(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], math.Cos)
}
func tan(vectors []Vector) Vector {
	return genericOpperation1V(vectors[0], math.Tan)
}

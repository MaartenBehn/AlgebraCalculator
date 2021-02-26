package V2

import (
	"math"
)

var mathFunctions = []Function{
	{"sqrt", sqrt, 1},
	{"degree", degree, 1},
	{"radians", radians, 1},
}

type Function struct {
	name            string
	function        func([]Vector) Vector
	attributeAmount int
}

func (f Function) getName() string {
	return f.name
}

func (f Function) getType() int {
	return TypFunction
}

func (f Function) getRank() int {
	return RankFunc
}

func (f Function) solve(term *Term, index int) {
	attributs := term.getSub(index+1, index+1+f.attributeAmount)

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

		term.setSub(index, index+1+f.attributeAmount,
			Term{parts: []TermPart{result}})
	}
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

func (f Function) print() {
	print(f.name)
}

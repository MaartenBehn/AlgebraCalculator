package V2

import (
	"fmt"
	"log"
	"math"
)

var mathOperators = []Operator{
	{",", appandVector, RankAppend},
	{"+", add, RankAddSub},
	{"-", sub, RankAddSub},
	{"*", mul, RankMul},
	{"/", div, RankMul},
	{"^", pow, RankMul},
	{"pow", pow, RankPow},
	{"dot", dot, RankFunc},
}

type Operator struct {
	name     string
	function func(Vector, Vector) Vector
	rank     int
}

func (o Operator) getName() string {
	return o.name
}
func (o Operator) getType() int {
	return TypOpperator
}
func (o Operator) getRank() int {
	return o.rank
}
func (o Operator) isSolvable() bool {
	return true
}
func (o Operator) solve(term *Term, index int) bool {
	term1 := term.parts[index-1]
	term2 := term.parts[index+1]

	if term1.getType() == TypVector && term2.getType() == TypVector {
		result := o.function(term1.(Vector), term2.(Vector))

		term.setSub(index-1, index+1, NewTerm([]ITermPart{result}))
	}
	return false
}
func (o Operator) print() {
	fmt.Print(o.name)
}

func appandVector(x Vector, y Vector) Vector {
	result := Vector{}
	result.append(x)
	result.append(y)
	return result
}
func add(x Vector, y Vector) Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 + f2
	})
}
func sub(x Vector, y Vector) Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 - f2
	})
}
func mul(x Vector, y Vector) Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 * f2
	})
}
func div(x Vector, y Vector) Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 / f2
	})
}
func pow(x Vector, y Vector) Vector {
	return genericOpperation2VScalar(x, y, math.Pow)
}
func dot(x Vector, y Vector) Vector {
	result := Vector{}

	if x.len == y.len {
		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[0] += x.values[i] * y.values[i]
		}

	} else {
		log.Panicf("Invalid vector Dimentions!")
	}

	return result
}

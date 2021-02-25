package AlgebraCalculator

import (
	"log"
	"math"
)

const maxRank int = 3

type Opperator struct {
	function    func(Term, int) Variable
	indexAcsess []int
	rank        int
}

var operators map[string]Opperator

func setUpOperatorMap() {
	operators = map[string]Opperator{}
	operators["+"] = Opperator{add, []int{-1, 1}, 1}
	operators["-"] = Opperator{sub, []int{-1, 1}, 1}
	operators["*"] = Opperator{mul, []int{-1, 1}, 2}
	operators["/"] = Opperator{div, []int{-1, 1}, 2}

	operators["^"] = Opperator{pow, []int{-1, 1}, 3}
	operators["pow"] = Opperator{pow, []int{-1, 1}, 3}
	operators["sqrt"] = Opperator{sqrt, []int{1}, 3}

	operators["abs"] = Opperator{abs, []int{1}, 3}
	operators["len"] = Opperator{magnitude, []int{1}, 3}
	operators["norm"] = Opperator{normalized, []int{1}, 3}
	operators["dist"] = Opperator{dist, []int{1}, 3}

	operators["dot"] = Opperator{dot, []int{-1, 1}, 3}

	operators["degree"] = Opperator{degree, []int{1}, 3}
	operators["radains"] = Opperator{radians, []int{1}, 3}

	operators["sin"] = Opperator{sin, []int{1}, 3}
	operators["asin"] = Opperator{asin, []int{1}, 3}
	operators["sinh"] = Opperator{sinh, []int{1}, 3}
	operators["asinh"] = Opperator{asinh, []int{1}, 3}

	operators["cos"] = Opperator{cos, []int{1}, 3}
	operators["acos"] = Opperator{acos, []int{1}, 3}
	operators["cosh"] = Opperator{cosh, []int{1}, 3}
	operators["acosh"] = Opperator{acosh, []int{1}, 3}

	operators["tan"] = Opperator{tan, []int{1}, 3}
	operators["atan"] = Opperator{atan, []int{1}, 3}
	operators["tanh"] = Opperator{tanh, []int{1}, 3}
	operators["atanh"] = Opperator{atanh, []int{1}, 3}
	operators["atan2"] = Opperator{atan2, []int{1}, 3}
}

func genericOpperation1(x Variable, opperation func(float64) float64) Variable {
	x.updateLen()

	result := Variable{}
	result.values = make([]float64, x.len)
	for i := 0; i < x.len; i++ {
		result.values[i] = opperation(x.values[i])
	}
	return result
}
func genericOpperation2(x Variable, y Variable, opperation func(float64, float64) float64) Variable {
	result := Variable{}

	x.updateLen()
	y.updateLen()

	if x.len == y.len {
		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[i] = opperation(x.values[i], y.values[i])
		}
	} else {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}
	return result
}
func genericOpperation2Scaler(x Variable, y Variable, opperation func(float64, float64) float64) Variable {
	result := Variable{}

	x.updateLen()
	y.updateLen()

	if x.len == y.len {

		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[i] = opperation(x.values[i], y.values[i])
		}

	} else if x.len == 1 {

		result.values = make([]float64, y.len)
		for i := 0; i < y.len; i++ {
			result.values[i] = opperation(x.values[0], y.values[i])
		}

	} else if y.len == 1 {

		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[i] = opperation(x.values[i], y.values[0])
		}

	} else {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}
	return result
}

func add(term Term, index int) Variable {
	return genericOpperation2Scaler(term.values[index-1], term.values[index+1],
		func(f1 float64, f2 float64) float64 {
			return f1 + f2
		})
}
func sub(term Term, index int) Variable {
	return genericOpperation2Scaler(term.values[index-1], term.values[index+1],
		func(f1 float64, f2 float64) float64 {
			return f1 - f2
		})
}
func mul(term Term, index int) Variable {
	return genericOpperation2Scaler(term.values[index-1], term.values[index+1],
		func(f1 float64, f2 float64) float64 {
			return f1 * f2
		})
}
func div(term Term, index int) Variable {
	return genericOpperation2Scaler(term.values[index-1], term.values[index+1],
		func(f1 float64, f2 float64) float64 {
			return f1 / f2
		})
}

func pow(term Term, index int) Variable {
	return genericOpperation2(term.values[index-1], term.values[index+1], math.Pow)
}
func sqrt(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Sqrt)
}

func abs(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Abs)
}
func magnitude(term Term, index int) Variable {
	input := term.values[index+1]
	input.updateLen()

	result := Variable{}
	result.values = make([]float64, 1)
	for i := 0; i < input.len; i++ {
		result.values[0] += math.Pow(input.values[i], 2)
	}
	result.values[0] = math.Sqrt(result.values[0])

	return result
}
func normalized(term Term, index int) Variable {
	input := term.values[index+1]
	input.updateLen()

	magnitude := magnitude(term, index)

	result := Variable{}
	result.values = make([]float64, input.len)
	for i := 0; i < input.len; i++ {
		result.values[i] = input.values[i] / magnitude.values[0]
	}

	return result
}
func dist(term Term, index int) Variable {
	input1 := term.values[index-1]
	input1.updateLen()

	input2 := term.values[index+1]
	input2.updateLen()

	result := Variable{}
	result.values = make([]float64, 1)

	if input1.len == 1 && input2.len == 1 {
		result.values = make([]float64, 1)
		result.values[0] = input1.values[0] - input2.values[0]

	} else if input1.len == input2.len {

		for i := 0; i < input1.len; i++ {
			result.values[0] = math.Pow(input1.values[i]-input2.values[i], 2)
		}
		result.values[0] = math.Sqrt(result.values[0])

	} else {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}

	return result
}

func dot(term Term, index int) Variable {
	input1 := term.values[index-1]
	input1.updateLen()

	input2 := term.values[index+1]
	input2.updateLen()

	result := Variable{}
	result.values = make([]float64, input1.len)
	if input1.len == input2.len {

		for i := 0; i < input1.len; i++ {
			result.values[0] += input1.values[i] * input2.values[i]
		}

	} else {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}

	return result
}

func degree(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], func(f float64) float64 {
		return f * (180.0 / math.Pi)
	})
}
func radians(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], func(f float64) float64 {
		return f * (math.Pi / 180.0)
	})
}

func sin(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Sin)
}
func asin(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Asin)
}
func sinh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Sinh)
}
func asinh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Asinh)
}

func cos(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Cos)
}
func acos(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Acos)
}
func cosh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Cosh)
}
func acosh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Acosh)
}

func tan(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Tan)
}
func atan(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Atan)
}
func tanh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Tanh)
}
func atanh(term Term, index int) Variable {
	return genericOpperation1(term.values[index+1], math.Atanh)
}
func atan2(term Term, index int) Variable {
	return genericOpperation2(term.values[index-1], term.values[index+1], math.Atan2)
}

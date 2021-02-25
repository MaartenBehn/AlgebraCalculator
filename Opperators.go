package AlgebraCalculator

import (
	"log"
	"math"
)

type Opperation struct {
	methode func(Variable, Variable) Variable
	before  bool
	after   bool
	rank    int
}

const maxRank int = 3

var operators map[string]Opperation

func setUpOperatorMap() {
	operators = map[string]Opperation{}

	operators["+"] = Opperation{add, true, true, 1}
	operators["-"] = Opperation{sub, true, true, 1}
	operators["*"] = Opperation{mul, true, true, 2}
	operators["/"] = Opperation{div, true, true, 2}

	operators["^"] = Opperation{pow, true, true, 3}
	operators["pow"] = Opperation{pow, true, true, 3}
	operators["sqrt"] = Opperation{sqrt, false, true, 3}

	operators["abs"] = Opperation{abs, false, true, 3}

	operators["len"] = Opperation{magnitude, false, true, 3}
	operators["norm"] = Opperation{normalized, false, true, 3}
	operators["dist"] = Opperation{dist, true, true, 3}

	operators["dot"] = Opperation{dot, true, true, 3}

	operators["degree"] = Opperation{degree, false, true, 3}
	operators["radians"] = Opperation{radians, false, true, 3}

	operators["sin"] = Opperation{sin, false, true, 3}
	operators["sinh"] = Opperation{sinh, false, true, 3}
	operators["asin"] = Opperation{asin, false, true, 3}
	operators["asinh"] = Opperation{asinh, false, true, 3}

	operators["cos"] = Opperation{cos, false, true, 3}
	operators["cosh"] = Opperation{cosh, false, true, 3}
	operators["acos"] = Opperation{acos, false, true, 3}
	operators["acosh"] = Opperation{acosh, false, true, 3}

	operators["tan"] = Opperation{tan, false, true, 3}
	operators["tanh"] = Opperation{tanh, false, true, 3}
	operators["atan"] = Opperation{atan, false, true, 3}
	operators["atanh"] = Opperation{atanh, false, true, 3}

	operators["atan2"] = Opperation{atan2, true, true, 3}
}

func basicOpperationWithScaler(x Variable, y Variable, opperation func(float64, float64) float64) Variable {
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
func basicMathFunction(x Variable, opperation func(float64) float64) Variable {
	x.updateLen()

	result := Variable{}
	result.values = make([]float64, x.len)
	for i := 0; i < x.len; i++ {
		result.values[i] = opperation(x.values[i])
	}
	return result
}

func add(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, func(f1 float64, f2 float64) float64 {
		return f1 + f2
	})
}
func sub(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, func(f1 float64, f2 float64) float64 {
		return f1 - f2
	})
}
func mul(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, func(f1 float64, f2 float64) float64 {
		return f1 * f2
	})
}
func div(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, func(f1 float64, f2 float64) float64 {
		return f1 / f2
	})
}

func pow(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, func(f1 float64, f2 float64) float64 {
		return math.Pow(f1, f2)
	})
}
func sqrt(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Sqrt)
}

func abs(x Variable, y Variable) Variable {
	result := Variable{}

	result.values = make([]float64, 1)

	y.updateLen()
	for i := 0; i < y.len; i++ {
		result.values[i] = math.Abs(y.values[i])
	}
	return result
}

func magnitude(x Variable, y Variable) Variable {

	result := Variable{}
	result.values = make([]float64, 1)

	y.updateLen()
	for i := 0; i < y.len; i++ {
		result.values[0] += math.Pow(y.values[i], 2)
	}
	result.values[0] = math.Sqrt(result.values[0])

	return result
}
func normalized(x Variable, y Variable) Variable {

	result := Variable{}

	magnitude := magnitude(Variable{}, y)

	result.values = make([]float64, y.len)
	for i := 0; i < y.len; i++ {
		result.values[i] = y.values[i] / magnitude.values[0]
	}

	return result
}
func dist(x Variable, y Variable) Variable {
	result := Variable{}

	x.updateLen()
	y.updateLen()

	if x.len == 1 && y.len == 1 {
		result.values = make([]float64, 1)
		result.values[0] = x.values[0] - y.values[0]

	} else if x.len == y.len {
		result = magnitude(x, sub(x, y))
	} else {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}

	return result
}

func dot(x Variable, y Variable) Variable {

	x.updateLen()
	y.updateLen()

	if x.len != y.len {
		log.Panicf("%s \nInvalid vector Dimentions!", currentLine)
	}

	result := Variable{}
	result.values = make([]float64, 1)
	for i := 0; i < x.len; i++ {
		result.values[0] += x.values[i] * y.values[i]
	}
	return result
}

func degree(x Variable, y Variable) Variable {
	return basicMathFunction(y, func(f float64) float64 {
		return f * (180.0 / math.Pi)
	})
}
func radians(x Variable, y Variable) Variable {
	return basicMathFunction(y, func(f float64) float64 {
		return f * (math.Pi / 180.0)
	})
}
func sin(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Sin)
}
func sinh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Sinh)
}
func asin(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Asin)
}
func asinh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Asinh)
}
func cos(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Cos)
}
func cosh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Cosh)
}
func acos(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Acos)
}
func acosh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Acosh)
}
func tan(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Tan)
}
func tanh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Tanh)
}
func atan(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Atan)
}
func atanh(x Variable, y Variable) Variable {
	return basicMathFunction(y, math.Atanh)
}
func atan2(x Variable, y Variable) Variable {
	return basicOpperationWithScaler(x, y, math.Atan2)
}

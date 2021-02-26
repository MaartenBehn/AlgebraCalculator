package V2

var mathOperators = []Operator{
	{"+", add, RankAddSub},
	{"-", sub, RankAddSub},
	{"*", mul, RankMul},
	{"/", div, RankMul},
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

func (o Operator) solve(term *Term, index int) {
	term1 := term.parts[index-1]
	term2 := term.parts[index+1]

	if term1.getType() == TypVector && term2.getType() == TypVector {
		result := o.function(term1.(Vector), term2.(Vector))

		term.setSub(index-1, index+1,
			Term{parts: []TermPart{result}})
	}
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

func (o Operator) print() {
	print(o.name)
}

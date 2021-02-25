package AlgebraCalculator

import (
	"fmt"
	"math"
)

type Variable struct {
	name   string
	values []float64
	len    int
}

func (v *Variable) updateLen() {
	v.len = len(v.values)
}

func (v Variable) getSubVariable(configuration string) (bool, Variable) {
	v.updateLen()

	subVariable := Variable{
		name: v.name + "." + configuration,
	}

	for _, indexChar := range configuration {
		index := (int(indexChar) - '0') - 1

		if index >= v.len {
			return true, subVariable
		}

		subVariable.values = append(subVariable.values, v.values[index])
	}

	return false, subVariable
}

func (v Variable) print() {

	fmt.Print(v.name + " : ")
	for _, value := range v.values {

		if value == math.Trunc(value) {
			fmt.Printf("%.0f ", value)
		} else {
			fmt.Printf("%.4f ", value)
		}

	}
	fmt.Print("\n")
}

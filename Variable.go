package AlgebraCalculator

import "fmt"

type Variable struct {
	name   string
	values []float64
	len    int
}

func (v *Variable) addValue(value Variable) {
	for _, number := range value.values {
		v.values = append(v.values, number)
	}
	v.updateLen()
}

func (v *Variable) updateLen() {
	v.len = len(v.values)
}

func (v *Variable) getSubVariable(configuration string) (bool, Variable) {

	subVariable := Variable{
		name: v.name + "." + configuration,
	}

	for _, indexChar := range configuration {
		index := (int(indexChar) - '0') - 1

		if index >= v.len {
			return false, subVariable
		}

		subVariable.values = append(subVariable.values, v.values[index])
	}

	return true, subVariable
}

func (v *Variable) print() {

	fmt.Print(v.name + " : ")
	for _, value := range v.values {
		fmt.Printf("%f ", value)
	}
	fmt.Print("\n")
}

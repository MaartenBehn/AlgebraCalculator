package V2

import "fmt"

type Variable struct {
	name string
}

func (v Variable) getName() string {
	return v.name
}
func (v Variable) getType() int {
	return TypVariable
}
func (v Variable) isSolvable() bool {
	return false
}
func (v Variable) print() {
	fmt.Print(v.name)
}
func (v Variable) getSimplify() int {
	return SimplifyVariable
}

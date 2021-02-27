package V2

import "fmt"

type Brace struct {
	opening bool
}

func (b Brace) getType() int {
	return TypBrace
}
func (b Brace) isSolvable() bool {
	return false
}
func (b Brace) print() {
	if b.opening {
		fmt.Print("(")
	} else {
		fmt.Print(")")
	}
}
func (b Brace) getSimplify() int {
	return SimplifyNone
}

package V2

import "fmt"

type SubOperation struct {
	configuration string
}

func (o SubOperation) getType() int {
	return TypSubOperation
}
func (o SubOperation) getRank() int {
	return RankSubOpperation
}
func (o SubOperation) isSolvable() bool {
	return true
}
func (o SubOperation) solve(term *Term, index int) bool {
	termPart := term.parts[index-1]

	if termPart.getType() == TypVector {

		vector := termPart.(Vector)

		if len(o.configuration) <= vector.len {

			result := Vector{}
			for _, indexChar := range o.configuration {
				index := (int(indexChar) - '0') - 1
				result.append(Vector{[]float64{vector.values[index]}, 1})
			}
			term.setSub(index-1, index, NewTerm([]ITermPart{result}))
		}
	}
	return false
}
func (o SubOperation) print() {
	fmt.Print("." + o.configuration)
}

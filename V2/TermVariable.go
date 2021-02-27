package V2

import "fmt"

var publicTerms []TermVariable

type TermVariable struct {
	Term
	name      string
	variables []Variable
}

func (t TermVariable) getType() int {
	return TypTermVariable
}
func (t TermVariable) isSolvable() bool {
	return true
}
func (t TermVariable) getRank() int {
	return RankTerm
}
func (t TermVariable) solve(term *Term, index int) bool {

	inputParts := make([]ITermPart, len(t.variables))
	for i := 0; i < len(t.variables); i++ {
		if len(term.parts) <= index+1+i || (term.parts[index+1+i].getType() != TypVector && term.parts[index+1+i].getType() != TypTermVariable) {
			return false
		}
		inputParts[i] = term.parts[index+1+i]
	}

	baseTerm := NewTerm(make([]ITermPart, len(t.parts)))
	for i, termPart := range t.parts {
		baseTerm.parts[i] = termPart
	}

	for i, variable := range t.variables {
		for j, termPart := range baseTerm.parts {
			if termPart.getType() == TypVariable && variable.name == termPart.(Variable).name {
				baseTerm.parts[j] = inputParts[i]
			}
		}
	}

	term.setSub(index, index+len(t.variables), baseTerm)

	return true
}
func (t TermVariable) print() {
	fmt.Printf("%s", t.name)

	if len(t.variables) != 0 {
		fmt.Print("<")
		for i, variable := range t.variables {
			fmt.Print(variable.name)

			if i < len(t.variables)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print(">")
	}
	fmt.Print(" : ")

	for _, termPart := range t.parts {
		termPart.print()
		fmt.Print(" ")
	}
}
func (t TermVariable) getSimplify() int {
	return SimplifyNone
}

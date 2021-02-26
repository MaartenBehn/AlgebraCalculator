package V2

type Term struct {
	parts []TermPart
}

func (term *Term) setSub(start int, end int, subTerm Term) {
	termVar := *term
	newParts := append(termVar.parts[:start], subTerm.parts...)
	term.parts = append(newParts, termVar.parts[end+1:]...)
}
func (term *Term) getSub(start int, end int) Term {
	return Term{term.parts[start : end+1]}
}

const (
	TypVector       = 1
	TypVariable     = 2
	TypOpperator    = 3
	TypFunction     = 4
	TypBrace        = 5
	TypSubOperation = 6
	TypTermVariable = 7
)

type TermPart interface {
	getType() int
	isSolvable() bool
	print()
}

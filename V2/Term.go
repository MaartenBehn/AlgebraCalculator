package V2

import "fmt"

type Term struct {
	parts []TermPart
}

func (t Term) print() {
	for _, termPart := range t.parts {
		termPart.print()
		fmt.Print(" ")
	}
}

func (term *Term) setSub(start int, end int, subTerm Term) {
	newParts := append(term.parts[:start-1], subTerm.parts...)
	term.parts = append(newParts, term.parts[end:]...)
}
func (term *Term) getSub(start int, end int) Term {
	return Term{term.parts[start:end]}
}

const (
	TypVector       int = 1
	TypVariable     int = 2
	TypOpperator    int = 3
	TypFunction     int = 4
	TypBrace        int = 5
	TypSubOperation int = 6
)

type TermPart interface {
	getType() int
	print()
}

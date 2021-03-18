package AlgebraCalculator

import (
	"AlgebraCalculator/log"
	"testing"
)

var testTermFunctions = []string{

	"a t = t + 2",
	"b x y = x * ( 1 , 2 ) + ( 8 , 7 ) + ( 1 , 1 ) * y",
	"c t = gauss a",
	"d t = gauss b",
}

func TestTermFunction(t *testing.T) {
	terms = nil

	for _, vectorTerm := range testTermFunctions {
		term, err := parseTerm(vectorTerm)
		if err != nil {
			t.Error(err)
			continue
		}
		simplifyRoot(term.root)
		log.PrintLog()

		log.Print("In: " + vectorTerm + "\nGot: ")
		term.print()
		log.Print("\n\n")
		log.PrintLog()

		terms = append(terms, term)
	}
}

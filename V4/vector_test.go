package V4

import (
	"AlgebraCalculator/log"
	"testing"
)

var testVectorTerms = []string{

	"a = 1 + 2 , 2 + 4",
	"a = ( 1 , 2 )",
	"a = ( 1 , 2 ) + ( 1 , 2 )",
	"a t = t * ( 1 , 2 )",
	"a t = ( 1 , 2 ) * t",
}

func TestVector(t *testing.T) {
	for _, vectorTerm := range testVectorTerms {
		term, err := parseTerm(vectorTerm)
		if err != nil {
			t.Error(err)
		}
		simplifyRoot(term.root)
		log.PrintLog()

		log.Print("In: " + vectorTerm + "\nGot: ")
		term.print()
		log.Print("\n\n")
		log.PrintLog()
	}
}

package AlgebraCalculator

import (
	"testing"
)

var testVectorTerms = []string{

	"a t = ( 1 + t , 2 * t , 4 ) dot ( t , 4 , 4 )",
	"a = 1 + 2 , 2 + 4",
	"a = ( 1 , 2 )",
	"a = ( 1 , 2 ) + ( 1 , 2 )",
	"a t = t * ( 1 , 2 )",
	"a t = ( 1 , 2 ) * t",
	"a = len ( 1 , 2 )",
	"a = ( -1 , -2 ) dist ( 1 , 2 )",
	"a = ( 1 , 2 ) . 1",
	"a = ( 1 , 2 ) . 2",
	"a = ( 1 , 2 ) . 12",
	"a = ( 1 , 2 ) . 21",
}

func TestVector(t *testing.T) {
	for _, vectorTerm := range testVectorTerms {
		term, err := parseTerm(vectorTerm)
		if err != nil {
			t.Error(err)
			continue
		}
		simplifyRoot(term.root)
		PrintLog()

		Print("In: " + vectorTerm + "\nGot: ")
		term.print()
		Print("\n\n")
		PrintLog()
	}
}

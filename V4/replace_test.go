package V4

import (
	"AlgebraCalculator/log"
	"testing"
)

var testTerms = []string{

	"a x y = ( x + y ) * ( x + y )",
	"a x y = ( x * y + x * y ) * ( x * y + x * y )",
	"a t = t + t",
	"a x y = ( x + x ) * ( x + x )",
	"a x y = x + 2 + y - x + 3 + y",
}

func TestReplace(t *testing.T) {
	for _, testTerm := range testTerms {
		term, err := parseTerm(testTerm)
		if err != nil {
			t.Error(err)
		}
		simplifyRoot(term.root)
		log.PrintLog()

		log.Print("In: " + testTerm + "\nGot: ")
		term.print()
		log.Print("\n\n")
		log.PrintLog()
	}

}

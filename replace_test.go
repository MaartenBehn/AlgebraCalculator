package AlgebraCalculator

import (
	"testing"
)

var testTerms = []string{

	"a x = deriv ( 2 * x pow 2 + 1 * x )",
	"a x = deriv ( 2 * x + 1 )",
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
		PrintLog()

		Print("In: " + testTerm + "\nGot: ")
		term.print()
		Print("\n\n")
		PrintLog()
	}
}

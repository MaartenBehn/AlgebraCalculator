package V4

import (
	"AlgebraCalculator/log"
	"testing"
)

var testPrints = []struct {
	in  string
	out string
}{
	{"a = 1", "a = 1"},
	{"a t = t", "a<t> = t"},
	{"a t = t + t", "a<t> = t + t"},
	{"a t = t + t * t", "a<t> = t + t * t"},
	{"a t = ( t + t ) * t", "a<t> = (t + t) * t"},
}

func TestPrint(t *testing.T) {
	for _, testPrint := range testPrints {
		term, err := parseTerm(testPrint.in)
		if err != nil {
			t.Error(err)
		}
		term.print()

		log := log.GetLog()
		if log != testPrint.out {
			t.Error("Got \"" + log + "\" but wanted \"" + testPrint.out + "\"")
		}
	}
}

package V1

import "testing"

func TestParseTerm(t *testing.T) {
	term, err := parseTerm("a = 4 + 4")
	if err != nil {
		t.Error(err)
	}

	if term.root.data != "+" {
		t.Error("Fail")
	}

	term, err = parseTerm("a t = t + 4")
	if err != nil {
		t.Error(err)
	}

	if term.root.data != "+" {
		t.Error("Fail")
	}
}

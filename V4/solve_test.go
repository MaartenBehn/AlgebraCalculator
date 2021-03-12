package V4

import "testing"

func TestSolve(t *testing.T) {
	Init()

	term, err := parseTerm("a = 4 + 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 8 {
		t.Error("Fail")
	}
}

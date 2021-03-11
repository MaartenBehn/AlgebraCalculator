package V4

import "testing"

func TestSimplifyRoot(t *testing.T) {

	term, err := parseTerm("a = 4 + 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 8 {
		t.Error("Fail")
	}

}

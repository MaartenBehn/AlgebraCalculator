package V4

import "testing"

func TestSolve(t *testing.T) {
	term, err := parseTerm("a = 4 + 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 8 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 - 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 0 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 * 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 16 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 / 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 1 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 + 4 + 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 12 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 + 4 - 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 4 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 + 4 * 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 20 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = 4 + 4 / 4")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 5 {
		t.Error("Fail")
	}

	term, err = parseTerm("a = sin 2")
	if err != nil {
		t.Error(err)
	}
	simplifyRoot(term.root)

	if term.root.dataNumber != 0.9092974268256816 {
		t.Error("Fail")
	}
}

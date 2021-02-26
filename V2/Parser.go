package V2

import "strings"

type NameBasedTermPart interface {
	TermPart
	getName() string
}

var nameBasedTermParts []NameBasedTermPart

func setUpNameBasedTermParts() {
	for _, x := range mathOperators {
		nameBasedTermParts = append(nameBasedTermParts, x)
	}

	for _, x := range mathFunctions {
		nameBasedTermParts = append(nameBasedTermParts, x)
	}
}

func parseTerm(text string) Term {
	term := Term{}

	parts := splitAny(text, " <>")
	parts = removeEmptiStrings(parts)

	for _, part := range parts {

		if part == "(" {
			term.parts = append(term.parts, Brace{true})
			continue
		}

		if part == ")" {
			term.parts = append(term.parts, Brace{false})
			continue
		}

		for _, nameBasedTermPart := range nameBasedTermParts {
			if part == nameBasedTermPart.getName() {
				term.parts = append(term.parts, nameBasedTermPart)
			}
		}

		if strings.Contains(part, ".") {
			parts2 := strings.Split(part, ".")
			parts2 = removeEmptiStrings(parts2)

			if len(parts2) != 2 {
				continue
			}

			if parts2[0] == ")" {
				term.parts = append(term.parts, Brace{false})
				term.parts = append(term.parts, SubOperation{parts2[1]})
				continue
			}

			if isNumber(parts2[0]) {
				term.parts = append(term.parts, getVector(part))
				continue
			}
		}

		if isNumber(part) {
			term.parts = append(term.parts, getVector(part))
			continue
		}
	}
	return term
}

package V2

import (
	"log"
	"strings"
)

type NameBasedTermPart interface {
	ITermPart
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

func parseTerm(text string) TermVariable {

	parts := strings.Split(text, ":")

	if len(parts) != 2 {
		log.Panic("Invalid Variable creation!")
	}
	parts1 := removeEmptiStrings(splitAny(parts[0], " <>"))

	var variables []Variable
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, Variable{name: parts1[i]})
	}

	term := Term{}

	parts2 := splitAny(parts[1], " <>")
	parts2 = removeEmptiStrings(parts2)

	for _, part := range parts2 {

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

		for _, publicTerm := range publicTerms {
			if part == publicTerm.name {
				term.parts = append(term.parts, publicTerm)
			}
		}

		for _, variable := range variables {
			if part == variable.name {
				term.parts = append(term.parts, variable)
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
	term.updateIndexes()

	publicTerm := TermVariable{
		Term:      term,
		name:      parts1[0],
		variables: variables,
	}

	return publicTerm
}

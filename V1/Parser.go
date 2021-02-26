package V1

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var lines []string
var currentLine string
var publicVariables []Variable

func Run(filename string) {

	setUpOperatorMap()

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(buf)

	lines = strings.Split(content, "\r\n")

	for _, line := range lines {
		if !strings.Contains(line, ":") {
			continue
		}
		currentLine = line

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Panicf("%s \nInvalid Variable creation!", currentLine)
		}
		parts1 := removeEmptiStrings(strings.Split(parts[0], " "))
		if len(parts1) != 1 {
			log.Panicf("%s \nInvalid Variable name! Too many Spaces.", currentLine)
		}
		var variable Variable

		variable.name = parts1[0]

		terms := strings.Split(parts[1], ",")
		for _, term := range terms {

			value, failed := calculateTerm(parseTerm(term))
			if failed {
				log.Panicf("%s \nInvalid Term.", currentLine)
			}

			variable.values = append(variable.values, value.values...)
		}

		publicVariables = append(publicVariables, variable)
	}

	for _, publicVariable := range publicVariables {
		publicVariable.print()
	}
}

func parseTerm(text string) Term {
	parts := splitAny(text, " <>")
	parts = removeEmptiStrings(parts)

	term := Term{
		values:      make([]Variable, len(parts)),
		opperations: make([]Opperator, len(parts)),
		subValue:    make([]string, len(parts)),
	}

	braceIndex := 0
	for i, part := range parts {

		breakLoop := false

		for key := range operators {
			if part == key {
				term.opperations[i] = operators[key]
				breakLoop = true
			}
		}
		if breakLoop {
			continue
		}

		if part == "(" {
			braceIndex++
			if braceIndex == 1 {
				term.openBraces = append(term.openBraces, i)
			}
			continue
		}

		if strings.Contains(part, ".") {
			parts2 := strings.Split(part, ".")

			if parts2[0] == ")" {
				braceIndex--
				if braceIndex == 0 {
					term.closeBraces = append(term.closeBraces, i)
				}
				term.subValue[i] = parts2[1]
				continue
			}

			if !isNumber(parts2[0]) {
				term.subValue[i] = parts2[1]
				term.values[i] = getPublicVariable(parts2[0])
				continue
			} else {
				term.values[i] = getNumber(part)
				continue
			}
		} else {

			if !isNumber(part) {
				term.values[i] = getPublicVariable(part)
				continue
			} else {
				term.values[i] = getNumber(part)
				continue
			}
		}
	}
	if braceIndex > 0 {
		log.Panicf("%s \nMissing closing brace", currentLine)
	} else if braceIndex > 0 {
		log.Panicf("%s \nMissing opening brace", currentLine)
	}

	return term
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func isNumber(text string) bool {
	is := true
	for _, char := range text {
		if !unicode.IsDigit(char) {
			is = false
		}
	}
	return is
}

func getPublicVariable(part string) Variable {
	value := Variable{}
	for _, publicVariable := range publicVariables {
		if part == publicVariable.name {
			value = publicVariable
		}
	}
	return value
}

func getNumber(part string) Variable {
	number, err := strconv.ParseFloat(part, 64)
	if err != nil {
		log.Panicf("%s \nCould not parse number!", currentLine)
	}
	return Variable{values: []float64{number}}
}

func removeEmptiStrings(strings []string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}

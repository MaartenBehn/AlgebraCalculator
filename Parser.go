package AlgebraCalculator

import (
	"io/ioutil"
	"log"
	"strings"
)

var lines []string
var currentLine string

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

			parts := strings.Split(term, " ")
			parts = removeEmptiStrings(parts)

			value, failed := calculateTerm(parts)
			if failed {
				log.Panicf("%s \nInvalid Term.", currentLine)
			}

			variable.addValue(value)
		}

		publicVariables = append(publicVariables, variable)
	}

	for _, publicVariable := range publicVariables {
		publicVariable.print()
	}
}

func removeEmptiStrings(strings []string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}

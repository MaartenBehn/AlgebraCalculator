package AlgebraCalculator

import (
	"AlgebraCalculator/log"
	"strings"
)

func Init(ruleStrings []string) {
	log.InitLog()
	initNamedNodeSlice()
	initRules(ruleStrings)
}

func Run(termsLines ...string) (results map[int]string, loggged string) {

	termMap := map[int]iNode{}
	for i, line := range termsLines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		term, err := parseTerm(line, termMap)
		if handelError(err) {
			continue
		}

		term.solveTerm()

		termMap[i] = term
	}
	loggged = log.GetLog()

	results = map[int]string{}
	for i, termNode := range termMap {
		termNode.(*term).printTerm()
		results[i] = log.GetLog()
	}

	return results, loggged
}

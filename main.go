package AlgebraCalculator

import (
	"AlgebraCalculator/log"
	"strings"
)

var rules [][]simpRule

func Init(ruleStrings []string) {
	log.InitLog()
	initNamedNodeSlice()

	for _, ruleString := range ruleStrings {
		rules = append(rules, parseRuleString(ruleString))
	}
}

func Run(termsTexts []string) (results []string, logged string) {

	termNodes := calculateTermList(termsTexts)
	logged = log.GetLog()

	for _, termNode := range termNodes {
		termNode.(*term).printTerm()
		results = append(results, log.GetLog())
	}

	return results, logged
}

func parseRuleString(ruleFile string) []simpRule {
	var simpRules []simpRule

	lines := removeEmptiStrings(splitAny(ruleFile, "\n\r"))

	for i, line := range lines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		rule, err := parseSimpRule(line)
		if handelError(err) {
			continue
		}

		rule.line = i
		simpRules = append(simpRules, rule)
	}
	return simpRules
}

func calculateTermList(lines []string) []iNode {

	var terms []iNode
	for _, line := range lines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		var term iNode
		term, err := parseTerm(line, terms)
		if handelError(err) {
			continue
		}

		term = solveTerm(term, rules)
		terms = append(terms, term)
	}
	return terms
}

func solveTerm(node iNode, rules [][]simpRule) iNode {
	term := node.(*term)

	log.Print("\n")
	term.printTerm()
	log.Print("\n")
	root := term.root

	log.Print("Parsed:")
	root.printTree(0)
	log.Print("\n")
	for _, ruleList := range rules {
		if ruleList == nil {
			continue
		}

		run := true
		for i := 0; i < 1000 && run; i++ {

			run = false
			run2 := true
			for run2 {
				run2 = root.sort()

				if run2 {
					log.Print("Sort:")
					root.printTree(0)
					log.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				run2 = root.solve()

				if run2 {
					log.Print("Solve:")
					root.printTree(0)
					log.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				for _, rule := range ruleList {
					run2 = rule.applyRule(root, rule.search)

					if run2 {
						log.Print("Simpify: ")
						log.Printf("%s", rule.base)
						root.printTree(0)
						log.Print("\n")

						root.solve()

						run = true
						break
					}
				}
			}
		}
	}
	term.root = root

	log.Print(" => ")
	term.printTerm()
	log.Print("\n")

	return term
}

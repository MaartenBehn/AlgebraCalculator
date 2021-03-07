package V3

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Run() {
	setUpNamedNodeSlice()

	var rules [][]SimpRule
	rules = append(rules, readRuleList("simpRulesExpand.txt"))
	rules = append(rules, readRuleList("simpRulesSumUp.txt"))

	terms := perfomTermList("input.txt", rules)

	for _, term := range terms {
		fmt.Print("\n")
		term.(*Term).printTerm()
	}
}

func readRuleList(path string) []SimpRule {
	var simpRules []SimpRule

	buf, err := ioutil.ReadFile(path)
	handelError(err)

	content := string(buf)
	lines := splitAny(content, "\n\r")

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

func perfomTermList(path string, rules [][]SimpRule) []INode {

	buf, err := ioutil.ReadFile(path)
	handelError(err)

	content := string(buf)
	lines := strings.Split(content, "\r\n")

	var terms []INode
	for _, line := range lines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		var term INode
		term, err := parseTerm(line, terms)
		if handelError(err) {
			continue
		}

		term = solveTerm(term, rules)
		terms = append(terms, term)
	}
	return terms
}

func solveTerm(term INode, rules [][]SimpRule) INode {
	fmt.Print("\n")
	term.(*Term).printTerm()
	fmt.Print("\n")
	root := term.(*Term).root

	fmt.Print("Parsed:")
	root.printTree(0)
	fmt.Print("\n")
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
					fmt.Print("Sort:")
					root.printTree(0)
					fmt.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				run2 = root.solve()

				if run2 {
					fmt.Print("Solve:")
					root.printTree(0)
					fmt.Print("\n")

					run = true
				}
			}

			run2 = true
			for run2 {
				for _, rule := range ruleList {
					run2 = rule.applyRule(root, rule.search)

					if run2 {
						fmt.Print("Simpify: ")
						fmt.Printf("%s", rule.base)
						root.printTree(0)
						fmt.Print("\n")

						root.solve()

						run = true
						break
					}
				}
			}
		}
	}
	term.(*Term).root = root

	fmt.Print(" => ")
	term.(*Term).printTerm()
	fmt.Print("\n")

	return term
}

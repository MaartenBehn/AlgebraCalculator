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

	terms := readTermList("input.txt")

	for _, term := range terms {
		solveTerm(term, rules)
	}
}

func readRuleList(path string) []SimpRule {
	var simpRules []SimpRule

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := splitAny(content, "\n\r")

	for i, line := range lines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		rule := parseSimpRule(line)
		rule.line = i

		simpRules = append(simpRules, rule)
	}
	return simpRules
}

func readTermList(path string) []INode {

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := strings.Split(content, "\r\n")

	var terms []INode
	for _, line := range lines {
		if !strings.Contains(line, "=") || strings.Contains(line, "//") {
			continue
		}

		term := parseTerm(line)
		terms = append(terms, term)
	}
	return terms
}

func solveTerm(term INode, rules [][]SimpRule) INode {
	term = term.copy()

	term.print()
	fmt.Print("\n")
	term.printTree(0)
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
				run2 = term.sort()

				if run2 {
					fmt.Print("Sort: ")
					term.print()
					fmt.Print("\n")

					run = true
				}
			}

			term.solve()

			run2 = true
			for run2 {
				for _, rule := range ruleList {
					run2 = rule.applyRule(term, rule.search)

					if run2 {
						fmt.Printf("%d %s \n", rule.line, rule.base)
						term.print()
						fmt.Print("\n")

						term.solve()

						run = true
						break
					}
				}
			}
		}

		fmt.Print("\n")
		term.printTree(0)
		fmt.Print("\n")
	}

	fmt.Print(" => ")
	term.print()
	fmt.Print("\n")

	return term
}

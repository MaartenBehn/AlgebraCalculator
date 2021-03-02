package V3

import (
	"io/ioutil"
	"strings"
)

func Run() {
	setUpNamedNodeSlice()

	buf, err := ioutil.ReadFile("simpRules.txt")
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := splitAny(content, "\n\r")

	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}

		rule := parseSimpRule(line)

		simpRules = append(simpRules, rule)
	}

	buf, err = ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content = string(buf)
	lines = strings.Split(content, "\r\n")

	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}

		publicTerm := parseTerm(line)

		result := publicTerm.copy()

		run := true
		for run {
			result.solve()

			run = false
			for _, rule := range simpRules {
				if rule.tryRule(result, rule.search) {
					run = true
				}
			}
		}

		result.print()
	}
}

package V2

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Run(filename string) {

	setUpNameBasedTermParts()

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := strings.Split(content, "\r\n")

	for _, line := range lines {
		if !strings.Contains(line, ":") {
			continue
		}

		publicTerm := parseTerm(line)

		publicTerm.Term = simplifyTermStep1(publicTerm.Term)

		publicTerm.Term = solveTerm(publicTerm.Term)

		publicTerms = append(publicTerms, publicTerm)

		publicTerm.print()
		fmt.Print("\n")
	}
}

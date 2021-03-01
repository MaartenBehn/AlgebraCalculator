package V3

import (
	"io/ioutil"
	"strings"
)

func Run(filename string) {

	setUpNamedNodeSlice()

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := strings.Split(content, "\r\n")

	for _, line := range lines {
		if !strings.Contains(line, "=") {
			continue
		}

		publicTerm := parseTerm(line)

		result := publicTerm.copy()

		result.solve()

		result.print()
	}
}

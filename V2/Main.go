package V2

import (
	"io/ioutil"
	"log"
	"strings"
)

func Run(filename string) {
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

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Panic("Invalid Variable creation!")
		}
		parts1 := removeEmptiStrings(strings.Split(parts[0], " "))
		if len(parts1) != 1 {
			log.Panic("Invalid Variable name! Too many Spaces.")
		}

		texts := strings.Split(parts[1], ",")
		for _, text := range texts {

			term := parseTerm(text)

			term = solveTerm(term)

			term.print()
		}
	}
}

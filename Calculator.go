package AlgebraCalculator

import (
	"log"
	"strconv"
)

var publicVariables []Variable

func calculateTerm(parts []string) (Variable, bool) {

	if len(parts) == 0 {
		return Variable{}, true
	}

	values := make([]Variable, len(parts))
	opperations := make([]Opperation, len(parts))

	hasBrace := false
	var openBraceIndexs []int
	var closeBraceIndexs []int
	braceIndex := 0

	for i, part := range parts {
		breakLoop := false

		if part == "(" {
			hasBrace = true
			braceIndex++
			if braceIndex == 1 {
				openBraceIndexs = append(openBraceIndexs, i)
			}
			continue

		} else if part == ")" {
			braceIndex--
			if braceIndex == 0 {
				closeBraceIndexs = append(closeBraceIndexs, i)
			}
			continue
		}

		for key := range operators {
			if part == key {
				opperations[i] = operators[key]
				breakLoop = true
			}
		}
		if breakLoop {
			continue
		}

		breakLoop = false
		for _, publicVariable := range publicVariables {
			if part == publicVariable.name {
				values[i] = publicVariable
				breakLoop = true
			}
		}
		if breakLoop {
			continue
		}

		number, err := strconv.ParseFloat(part, 64)
		if err != nil {
			log.Panicf("%s \nCould not parse number!", currentLine)
		}
		values[i] = Variable{values: []float64{number}}
	}
	if braceIndex > 0 {
		log.Panicf("%s \nMissing closing brace", currentLine)
	} else if braceIndex > 0 {
		log.Panicf("%s \nMissing opening brace", currentLine)
	}

	if len(values) == 1 && len(values[0].values) > 0 {
		return values[0], false
	}

	if hasBrace {

		braceLengthSum := 0
		for i, openIndex := range openBraceIndexs {
			closeIndex := closeBraceIndexs[i]

			braceParts := parts[openIndex+1 : closeIndex]
			value, failed := calculateTerm(braceParts)

			if failed {
				return Variable{}, false
			}

			values = append(values[:openIndex-braceLengthSum], values[closeIndex-braceLengthSum:]...)
			values[openIndex-braceLengthSum] = value

			opperations = append(opperations[:openIndex-braceLengthSum], opperations[closeIndex-braceLengthSum:]...)

			braceLengthSum += closeIndex - openIndex
		}
	}
	var opperationsExecutionOrder []int

	for i := maxRank; i >= 1; i-- {
		for j, opperation := range opperations {
			if opperation.rank == i {
				opperationsExecutionOrder = append(opperationsExecutionOrder, j)
			}
		}
	}

	copyValueFrom := make([]int, len(values))
	for i := range copyValueFrom {
		copyValueFrom[i] = -2
	}

	var lastOpperationValue int
	for _, id := range opperationsExecutionOrder {

		beforeValue := Variable{}
		var beforeIndex int
		if opperations[id].before {
			beforeIndex = id - 1
			for copyValueFrom[beforeIndex] != -2 {
				beforeIndex = copyValueFrom[beforeIndex]
			}
			beforeValue = values[beforeIndex]
		}

		afterValue := Variable{}
		var afterIndex int
		if opperations[id].after {
			afterIndex = id + 1
			for copyValueFrom[afterIndex] != -2 {
				afterIndex = copyValueFrom[afterIndex]
			}
			afterValue = values[afterIndex]
		}

		value := opperations[id].methode(beforeValue, afterValue)

		values[id] = value

		if opperations[id].before {
			copyValueFrom[beforeIndex] = id
		}

		if opperations[id].after {
			copyValueFrom[afterIndex] = id
		}

		lastOpperationValue = id
	}

	return values[lastOpperationValue], false
}

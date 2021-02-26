package V1

type Term struct {
	values      []Variable
	opperations []Opperator
	subValue    []string
	openBraces  []int
	closeBraces []int
}

func (term Term) subTerm(start int, end int) Term {
	subTerm := Term{
		values:      term.values[start+1 : end],
		opperations: term.opperations[start+1 : end],
		subValue:    term.subValue[start+1 : end],
	}

	for i, openBrace := range term.openBraces {
		if openBrace > start && openBrace < end {
			subTerm.openBraces = append(subTerm.openBraces, openBrace-start-1)
			subTerm.closeBraces = append(subTerm.closeBraces, term.closeBraces[i]-start-1)
		}
	}

	return subTerm
}

func calculateTerm(term Term) (Variable, bool) {

	// Braces
	if len(term.openBraces) > 0 {
		braceLengthSum := 0
		for i, openIndex := range term.openBraces {
			closeIndex := term.closeBraces[i]

			value, failed := calculateTerm(term.subTerm(openIndex, closeIndex))

			if failed {
				return Variable{}, true
			}

			term.values = append(term.values[:openIndex-braceLengthSum], term.values[closeIndex-braceLengthSum:]...)
			term.values[openIndex-braceLengthSum] = value

			term.opperations = append(term.opperations[:openIndex-braceLengthSum], term.opperations[closeIndex-braceLengthSum:]...)

			term.subValue = append(term.subValue[:openIndex-braceLengthSum], term.subValue[closeIndex-braceLengthSum:]...)

			braceLengthSum += closeIndex - openIndex
		}
	}

	for i, configuration := range term.subValue {
		if configuration == "" {
			continue
		}

		failed := false
		failed, term.values[i] = term.values[i].getSubVariable(configuration)

		if failed {
			return Variable{}, true
		}
	}

	// Execution order
	var opperationsExecutionOrder []int
	for i := maxRank; i >= 1; i-- {
		for j, opperation := range term.opperations {
			if opperation.rank == i {
				opperationsExecutionOrder = append(opperationsExecutionOrder, j)
			}
		}
	}

	// The actual calculations
	lastOpperationValue := 0
	for _, index := range opperationsExecutionOrder {

		value := term.opperations[index].function(term, index)
		term.values[index] = value

		for _, indexOffset := range term.opperations[index].indexAcsess {
			term.values[index+indexOffset] = value
		}

		lastOpperationValue = index
	}

	return term.values[lastOpperationValue], false

}

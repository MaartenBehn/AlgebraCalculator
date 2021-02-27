package V2

const (
	SimplifyOne      = 1
	SimplifyMinusOne = 2
	SimplifyTwo      = 3
	SimplifyAdd      = 4
	SimplifySub      = 5
	SimplifyMul      = 6
	SimplifyDiv      = 7
	SimplifyPow      = 8
	SimplifyVector   = 9
	SimplifyVariable = 10
	SimplifyNone     = 12
)

var simplifyRules = []SimplifyRule{
	{[]int{SimplifySub, SimplifyVector}, []int{SimplifyAdd, SimplifyMinusOne, SimplifyMul, SimplifyVector}},
	{[]int{SimplifySub, SimplifyVariable}, []int{SimplifyAdd, SimplifyMinusOne, SimplifyMul, SimplifyVariable}},

	{[]int{SimplifyVariable, SimplifyDiv, SimplifyVector}, []int{SimplifyOne, SimplifyDiv, SimplifyVector, SimplifyMul, SimplifyVariable}},
	{[]int{SimplifyVector, SimplifyDiv, SimplifyVariable}, []int{SimplifyVector, SimplifyMul, SimplifyVariable, SimplifyPow, SimplifyMinusOne}},
	{[]int{SimplifyVariable, SimplifyDiv, SimplifyVariable}, []int{}},

	{[]int{SimplifyVariable, SimplifyMul, SimplifyVector}, []int{SimplifyVector, SimplifyMul, SimplifyVariable}},
	{[]int{SimplifyVariable, SimplifyMul, SimplifyVariable}, []int{SimplifyVariable, SimplifyPow, SimplifyTwo}},
}

type SimplifyRule struct {
	search  []int
	replace []int
}

func simplifyTermStep1(term Term) Term {

	var foundIndex int
	for foundIndex != -1 {

		stop := false
		foundIndex = -1
		var foundRule SimplifyRule
		for _, rule := range simplifyRules {

			for i := range term.parts {

				checkIndex := 0
				varibleName := ""
				for j := i; j < len(term.parts); j++ {

					termPart := term.parts[j]
					if termPart.getSimplify() == rule.search[checkIndex] {
						checkIndex++

						if termPart.getSimplify() == SimplifyVariable && varibleName == "" {
							varibleName = termPart.(NameBasedTermPart).getName()
						} else if termPart.getSimplify() == SimplifyVariable && varibleName != termPart.(NameBasedTermPart).getName() {
							break
						}

						if checkIndex >= len(rule.search) {
							foundIndex = j - checkIndex + 1
							foundRule = rule
							stop = true
							break
						}
					} else {
						break
					}
				}
				if stop {
					break
				}
			}
			if stop {
				break
			}
		}

		if foundIndex == -1 {
			continue
		}

		replaceTerm := Term{}
		for _, rulePart := range foundRule.replace {

			var termPart ITermPart
			switch rulePart {
			case SimplifyOne:
				termPart = Vector{
					values: []float64{1},
					len:    1,
				}
				break
			case SimplifyMinusOne:
				termPart = Vector{
					values: []float64{-1},
					len:    1,
				}
				break
			case SimplifyTwo:
				termPart = Vector{
					values: []float64{2},
					len:    1,
				}
				break
			case SimplifyAdd:
				termPart = findOperator("+")
				break
			case SimplifySub:
				termPart = findOperator("-")
				break
			case SimplifyMul:
				termPart = findOperator("*")
				break
			case SimplifyDiv:
				termPart = findOperator("/")
				break
			case SimplifyPow:
				termPart = findOperator("pow")
				break
			default:
				for i := foundIndex; i < foundIndex+len(foundRule.search); i++ {
					if term.parts[i].getSimplify() == rulePart {
						termPart = term.parts[i]
						break
					}
				}
				break
			}

			replaceTerm.parts = append(replaceTerm.parts, termPart)
		}

		term.setSub(foundIndex, foundIndex+len(foundRule.search)-1, replaceTerm)
	}

	var subTerms []Term
	startIndex := 0
	for i, termParts := range term.parts {
		if termParts.getType() == TypOpperator && termParts.(NameBasedTermPart).getName() == "+" && startIndex < i {
			subTerms = append(subTerms, term.getSub(startIndex, i-1))
			startIndex = i + 1
		}
	}
	subTerms = append(subTerms, term.getSub(startIndex, len(term.parts)-1))

	var alreadyMerged []int
	var mergedTerms []Term
	for i, subTerm := range subTerms {

		merged := false
		for _, index := range alreadyMerged {
			if i == index {
				merged = true
				break
			}
		}
		if merged {
			continue
		}

		var variables []NameBasedTermPart
		for _, termPart := range subTerm.parts {
			if termPart.getSimplify() == SimplifyVariable {
				variables = append(variables, termPart.(NameBasedTermPart))
			}
		}

		var sameSubTerms []Term
		for j, testSubTerm := range subTerms {

			if i == j {
				continue
			}

			var testVariables []NameBasedTermPart
			for _, termPart := range testSubTerm.parts {
				if termPart.getSimplify() == SimplifyVariable {
					testVariables = append(testVariables, termPart.(NameBasedTermPart))
				}
			}

			var counter int
			if len(variables) > len(testVariables) {
				counter = len(variables)
			} else {
				counter = len(testVariables)
			}

			for _, variable := range variables {
				for _, testVariable := range testVariables {
					if variable.getName() == testVariable.getName() {
						counter--
						break
					}
				}
			}

			if counter == 0 {
				sameSubTerms = append(sameSubTerms, testSubTerm)
				alreadyMerged = append(alreadyMerged, j)
			}
		}

		for _, sameSubTerm := range sameSubTerms {

			var vectors []ITermPart
			for _, termPart := range sameSubTerm.parts {
				if termPart.getSimplify() == SimplifyVector {
					vectors = append(vectors, termPart)
				}
			}

			if len(vectors) == 0 {
				vectors = append(vectors, Vector{
					values: []float64{1},
					len:    1,
				})
			}

			for i, vector := range vectors {
				var addTerm Term
				if i == 0 {
					addTerm = NewTerm([]ITermPart{vector, findOperator("+")})
				} else {
					addTerm = NewTerm([]ITermPart{vector, findOperator("*")})
				}

				subTerm.parts = append(addTerm.parts, subTerm.parts...)
			}

			subTerm.updateIndexes()

		}

		mergedTerms = append(mergedTerms, subTerm)
	}

	var finalParts []ITermPart
	for i, mergedTerm := range mergedTerms {
		if i < len(mergedTerms)-1 {
			mergedTerm.parts = append(mergedTerm.parts, findOperator("+"))
		}
		finalParts = append(finalParts, mergedTerm.parts...)
	}

	return NewTerm(finalParts)
}

package V2

const (
	SimplifyAdd      = -100
	SimplifySub      = -101
	SimplifyMul      = -102
	SimplifyDiv      = -103
	SimplifyPow      = -104
	SimplifyVector   = -105
	SimplifyVariable = -106
	SimplifyFunction = -107
	SimplifyNone     = -108
)

var simplifyRules = []SimplifyRule{

	{[]float64{SimplifySub, SimplifyVector}, []float64{SimplifyAdd, -1, SimplifyMul, SimplifyVector}},
	{[]float64{SimplifySub, SimplifyVariable}, []float64{SimplifyAdd, -1, SimplifyMul, SimplifyVariable}},

	{[]float64{SimplifyVariable, SimplifyDiv, SimplifyVector}, []float64{1.0, SimplifyDiv, SimplifyVector, SimplifyMul, SimplifyVariable}},
	{[]float64{SimplifyVariable, SimplifyDiv, SimplifyVariable}, []float64{1}},

	{[]float64{SimplifyVariable, SimplifyMul, SimplifyVector}, []float64{SimplifyVector, SimplifyMul, SimplifyVariable}},

	{[]float64{SimplifyVariable, SimplifyPow, 1}, []float64{SimplifyVariable}},
	{[]float64{SimplifyVariable, SimplifyPow, 2}, []float64{SimplifyVariable, SimplifyMul, SimplifyVariable}},
	{[]float64{SimplifyVariable, SimplifyPow, 3}, []float64{SimplifyVariable, SimplifyMul, SimplifyVariable, SimplifyMul, SimplifyVariable}},
	{[]float64{SimplifyVariable, SimplifyPow, 4}, []float64{SimplifyVariable, SimplifyMul, SimplifyVariable, SimplifyMul, SimplifyVariable, SimplifyMul, SimplifyVariable}},
}

type SimplifyRule struct {
	search  []float64
	replace []float64
}

func simplifyTerm(term Term) Term {

	term = simplifyTermStep2(term)
	term = simplifyTermStep3(term)

	return term
}

func simplifyTermStep1(term Term) Term {

	return term
}

func simplifyTermStep2(term Term) Term {

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
					if termPart.getSimplify() == rule.search[checkIndex] ||
						(termPart.getSimplify() == SimplifyVector && termPart.(Vector).len == 1 && termPart.(Vector).values[0] == rule.search[checkIndex]) {
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
				found := false
				for i := foundIndex; i < foundIndex+len(foundRule.search); i++ {
					if term.parts[i].getSimplify() == rulePart {
						termPart = term.parts[i]
						found = true
						break
					}
				}
				if found {
					break
				}

				termPart = Vector{
					values: []float64{rulePart},
					len:    1,
				}

				break
			}

			replaceTerm.parts = append(replaceTerm.parts, termPart)
		}

		term.setSub(foundIndex, foundIndex+len(foundRule.search)-1, replaceTerm)
	}

	return term
}

type SimplifyExpression struct {
	baseTerm      Term
	variableTerms []Term
	exponens      []float64
	vectorTerms   []Term
	failed        bool
	merged        bool
}

func simplifyTermStep3(term Term) Term {

	var subTerms []Term
	startIndex := 0
	for i, termParts := range term.parts {
		if termParts.getType() == TypOpperator && termParts.(NameBasedTermPart).getName() == "+" && startIndex < i {
			subTerms = append(subTerms, term.getSub(startIndex, i-1))
			startIndex = i + 1
		}
	}
	subTerms = append(subTerms, term.getSub(startIndex, len(term.parts)-1))

	var expressions []SimplifyExpression

	for _, subTerm := range subTerms {

		expression := SimplifyExpression{}
		expression.baseTerm = subTerm
		for i, termPart := range subTerm.parts {

			if termPart.getSimplify() == SimplifyNone {
				expression.failed = true
				break
			} else if termPart.getSimplify() == SimplifyVariable || termPart.getSimplify() == SimplifyFunction {

				var variableTerm Term
				if termPart.getSimplify() == SimplifyVariable {
					variableTerm = NewTerm([]ITermPart{termPart})
				} else {
					variableTerm = subTerm.getSub(i, i+termPart.(MathFunction).attributeAmount)
				}

				index := -1
				for i, testvariableTerm := range expression.variableTerms {
					if areNameBasedTermsEqual(variableTerm, testvariableTerm) {
						index = i
					}
				}

				if i == 0 || subTerm.parts[i-1].getSimplify() == SimplifyMul {

					if index == -1 {
						index = len(expression.variableTerms)
						expression.variableTerms = append(expression.variableTerms, variableTerm)
						expression.exponens = append(expression.exponens, 0)
					}

					expression.exponens[index] += 1
				} else if subTerm.parts[i-1].getSimplify() == SimplifyDiv {

					if index == -1 {
						index = len(expression.variableTerms)
						expression.variableTerms = append(expression.variableTerms, variableTerm)
						expression.exponens = append(expression.exponens, 0)
					}

					expression.exponens[index] -= 1
				}

			} else if termPart.getSimplify() == SimplifyVector {

				if i == 0 || subTerm.parts[i-1].getSimplify() == SimplifyMul {
					expression.vectorTerms = append(expression.vectorTerms, NewTerm([]ITermPart{termPart}))
				} else if subTerm.parts[i-1].getSimplify() == SimplifyDiv {
					expression.vectorTerms = append(expression.vectorTerms, NewTerm([]ITermPart{Vector{[]float64{1}, 1}, subTerm.parts[i-1], termPart}))
				}
			}
		}

		if len(expression.variableTerms) > 0 && len(expression.vectorTerms) == 0 {
			expression.vectorTerms = append(expression.vectorTerms, NewTerm([]ITermPart{Vector{[]float64{1}, 1}}))
		}

		expressions = append(expressions, expression)
	}

	for i, expression := range expressions {
		if expression.failed || expression.merged {
			continue
		}

		var newTermParts []ITermPart

		newTermParts = append(newTermParts, Brace{true})

		for o, vectorTerm := range expression.vectorTerms {
			newTermParts = append(newTermParts, vectorTerm.parts...)

			if o < len(expression.vectorTerms)-1 {
				newTermParts = append(newTermParts, findOperator("*"))
			} else {
				newTermParts = append(newTermParts, findOperator("+"))
			}
		}

		for j, testExpression := range expressions {
			if i == j || expression.failed || expression.merged {
				continue
			}

			if len(expression.variableTerms) != len(testExpression.variableTerms) {
				continue
			}

			counter := len(expression.variableTerms)
			for i, variableTerm := range expression.variableTerms {
				for j, testVariableTerm := range testExpression.variableTerms {
					if areNameBasedTermsEqual(variableTerm, testVariableTerm) && expression.exponens[i] == testExpression.exponens[j] {
						counter--
					}
				}
			}

			if counter != 0 {
				continue
			}

			for o, vectorTerm := range testExpression.vectorTerms {
				newTermParts = append(newTermParts, vectorTerm.parts...)

				if o < len(testExpression.vectorTerms)-1 {
					newTermParts = append(newTermParts, findOperator("*"))
				} else {
					newTermParts = append(newTermParts, findOperator("+"))
				}
			}

			expressions[j].merged = true
		}

		newTermParts[len(newTermParts)-1] = Brace{false}

		if len(expression.variableTerms) > 0 {
			newTermParts = append(newTermParts, findOperator("*"))

			for o, variable := range expression.variableTerms {
				newTermParts = append(newTermParts, variable.parts...)

				if expression.exponens[o] != 1 {
					newTermParts = append(newTermParts, findOperator("pow"))
					newTermParts = append(newTermParts, Vector{[]float64{expression.exponens[o]}, 1})
				}

				if o < len(expression.variableTerms)-1 {
					newTermParts = append(newTermParts, findOperator("*"))
				}
			}
		}

		expressions[i].baseTerm = NewTerm(newTermParts)
	}

	var newTermParts []ITermPart

	for _, expression := range expressions {
		if !expression.failed && expression.merged {
			continue
		}

		if len(newTermParts) > 0 {
			newTermParts = append(newTermParts, findOperator("+"))
		}

		newTermParts = append(newTermParts, expression.baseTerm.parts...)
	}

	term = NewTerm(newTermParts)

	return term
}

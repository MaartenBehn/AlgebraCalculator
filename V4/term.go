package V4

import "AlgebraCalculator/log"

var parseTermFuncs []func(text string) *parserNode

func initTerm() {
	parseTermFuncs = append(parseTermFuncs,
		func(text string) *parserNode { return tryParseNumber(text) },

		func(text string) *parserNode { return tryParseOperator2(text, "+", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "-", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "*", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "/", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "pow", rankPow) },
		func(text string) *parserNode { return tryParseOperator2(text, ",", rankAppend) },

		func(text string) *parserNode { return tryParseOperator1(text, "sin", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "sinh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "asin", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "asinh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "cos", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "cosh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "acos", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "acosh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "tan", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "tanh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "atan", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator1(text, "atanh", rankMathFunction) },
		func(text string) *parserNode { return tryParseOperator2(text, "atan2", rankMathFunction) },

		func(text string) *parserNode { return tryParseVaraible(text) },
	)
}

type term struct {
	name      string
	variables []*node
	root      *node
}

func newTerm(name string) *term {
	return &term{name: name}
}
func (t *term) print() {
	log.Print(t.name)
	if len(t.variables) > 0 {

		log.Print("<")
		for i, variable := range t.variables {
			variable.print()
			if i < len(t.variables)-1 {
				log.Print(" ")
			}
		}
		log.Print(">")
	}
	log.Print(" = ")

	t.root.print()
}

func parseTerm(text string) (*term, error) {
	parts := splitAny(text, "=")

	if len(parts) != 2 {
		if len(parts) != 2 {
			return nil, newError(errorTypParsing, errorCriticalLevelPartial, "There is not a valid part before or after \"=\" in term!")
		}
	}

	parts1 := splitAny(parts[0], " <>")
	if len(parts1) == 0 {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "Term has no name!")
	}

	var variables []*node
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, newNode(parts1[i], 0, flagData, flagVariable))
	}

	parts2 := splitAny(parts[1], " <>")
	if len(parts2) == 0 {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "No Expression after \"=\"")
	}

	currentVariables = variables
	root, _, err := parseRoot(parseTermFuncs, parts2...)
	if handelError(err) {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "term could not be parsed!")
	}

	log.Print("Parse: \n")
	root.print()
	log.Print("\n")

	t := newTerm(parts1[0])
	t.variables = variables
	t.root = root.node
	return t, nil
}

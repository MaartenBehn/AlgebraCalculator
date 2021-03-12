package V4

var parseTermFuncs []func(text string) *parserNode

func initTerm() {
	parseTermFuncs = append(parseTermFuncs,
		func(text string) *parserNode { return tryParseNumber(text) },
		func(text string) *parserNode { return tryParseVaraible(text) },
		func(text string) *parserNode { return tryParseOperator2(text, "+", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "-", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "*", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "/", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "pow", rankPow) },
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
func parseTerm(text string) (*term, error) {
	parts := splitAny(text, "=")

	if len(parts) != 2 {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "There is not a valid part before or after \"=\" in term!")
	}

	parts1 := splitAny(parts[0], " <>")
	var variables []*node
	for i := 1; i < len(parts1); i++ {
		variables = append(variables, NewNode(parts1[i], 0, flagData, flagVariable))
	}

	parts2 := splitAny(parts[1], " <>")
	currentVariables = variables
	root, _, err := parseRoot(parseTermFuncs, parts2...)
	if handelError(err) {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "term could not be parsed!")
	}

	t := newTerm(parts1[0])
	t.variables = variables
	t.root = root.node
	return t, nil
}

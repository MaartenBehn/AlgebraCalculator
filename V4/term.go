package V4

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
		variables = append(variables, NewNode(parts1[i], flagData, flagVariable))
	}

	parts2 := splitAny(parts[1], " <>")
	root, _, err := parseRoot(parts2...)
	if handelError(err) {
		return nil, newError(errorTypParsing, errorCriticalLevelPartial, "term could not be parsed!")
	}

	t := newTerm(parts1[0])
	t.variables = variables
	t.root = root.node
	return t, nil
}

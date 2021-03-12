package V4

var parseReplaceFuncs []func(text string) *parserNode

func initReplace() {
	parseReplaceFuncs = append(parseReplaceFuncs,
		func(text string) *parserNode { return tryParseNumber(text) },
		func(text string) *parserNode { return tryParseReplaceRulePart(text) },
		func(text string) *parserNode { return tryParseOperator2(text, "+", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "-", rankAddSub) },
		func(text string) *parserNode { return tryParseOperator2(text, "*", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "/", rankMul) },
		func(text string) *parserNode { return tryParseOperator2(text, "pow", rankPow) },
	)

	ruleStrings := []string{
		"data_0 + data_0 = 2 * data_0",
		"( data_0 + data_1i ) * ( data_0 + data_1i ) = data_0 pow 2 + data_1 pow 2 + 2 * data_0 * data_1",
	}

	for _, ruleString := range ruleStrings {
		simpPatterns = append(simpPatterns, replace(parseReplaceRule(ruleString)))
	}
}

func parseReplaceRule(rule string) (*node, *node) {
	parts := splitAny(rule, "=")
	if len(parts) != 2 {
		handelError(newError(errorTypParsing, errorCriticalLevelFatal, "Rule has not two parts!"))
	}

	root0, _, err := parseRoot(parseReplaceFuncs, splitAny(parts[0], " ")...)
	if handelError(err) {
		handelError(newError(errorTypParsing, errorCriticalLevelFatal, "Rule could not be parsed!"))
	}

	root1, _, err := parseRoot(parseReplaceFuncs, splitAny(parts[1], " ")...)
	if handelError(err) {
		handelError(newError(errorTypParsing, errorCriticalLevelFatal, "Rule could not be parsed!"))
	}

	return root0.node, root1.node
}

type replaceDataBuffer []node

func (dataBuffer *replaceDataBuffer) checkAndSet(current *node, search *node) bool {
	index := len(*dataBuffer)
	for i, testnode := range *dataBuffer {
		if current.equal(&testnode) {
			index = i
			break
		}
	}

	id := int(search.data[0]) - '0'
	if id >= len(*dataBuffer) {
		*dataBuffer = append(*dataBuffer, *current)
	}

	debug := id == index || ((len(search.data) == 2 && search.data[1] == 'i') && id >= index)
	return debug
}

func replace(search *node, replace *node) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return checkReplace(root, search, &replaceDataBuffer{})
		},
		func(root *node) *node {
			dataBuffer := &replaceDataBuffer{}
			setReplaceDataBuffer(root, search, dataBuffer)

			newRoot := replace.copy()
			replaceNodes(newRoot, replace, dataBuffer)
			return newRoot
		},
	}
}
func checkReplace(current *node, search *node, dataBuffer *replaceDataBuffer) bool {
	if len(search.childs) != len(current.childs) {
		return false
	}
	for i, child := range current.childs {
		if !checkReplace(child, search.childs[i], dataBuffer) {
			return false
		}
	}

	if search.hasFlag(flagRulePart) {
		for flag, flagValue := range search.flagValues {
			if flagValue && !current.hasFlag(flag) && flag != flagRulePart {
				return false
			}
		}
		if !dataBuffer.checkAndSet(current, search) {
			return false
		}

	} else {
		if !current.equal(search) {
			return false
		}
	}

	return true
}
func setReplaceDataBuffer(current *node, search *node, dataBuffer *replaceDataBuffer) {
	for i, child := range current.childs {
		setReplaceDataBuffer(child, search.childs[i], dataBuffer)
	}

	if search.hasFlag(flagRulePart) {
		dataBuffer.checkAndSet(current, search)
	}
}
func replaceNodes(current *node, replacement *node, dataBuffer *replaceDataBuffer) {
	for i, child := range current.childs {
		replaceNodes(child, replacement.childs[i], dataBuffer)
	}

	childs := current.childs
	if replacement.hasFlag(flagRulePart) {
		id := int(replacement.dataNumber)

		*current = (*dataBuffer)[id]
	} else {
		*current = *replacement
	}
	current.childs = childs
}

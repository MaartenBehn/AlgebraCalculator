package V4

import "strconv"

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

		// Fix wrong Brace order
		"all_0i + ( all_1i + all_2i ) = all_0 + all_1 + all_2",
		"all_0i * ( all_1i * all_2i ) = all_0 * all_1 * all_2",

		// Brackte Rules:
		"( all_0 + all_0 ) * ( all_0 + all_0 ) = 4 * all_0 pow 2",

		// - Two
		"( all_0 + all_1 ) pow 2 = ( all_0 + all_1 ) * ( all_0 + all_1 )",
		"( all_0 - all_1 ) pow 2 = ( all_0 - all_1 ) * ( all_0 - all_1 )",

		"( all_0i + all_1i ) * ( all_0i + all_1i ) = all_0 pow 2 + all_1 pow 2 + 2 * all_0 * all_1",
		"( all_0i + all_1i ) * ( all_1i + all_0i ) = all_0 pow 2 + all_1 pow 2 + 2 * all_0 * all_1",

		"( all_0i + all_1i ) * ( all_0i - all_1i ) = all_0 pow 2 - all_1 pow 2",
		"( all_0i + all_1i ) * ( all_1i - all_0i ) = all_1 pow 2 - all_0 pow 2",

		"( all_0i - all_1i ) * ( all_0i + all_1i ) = all_0 pow 2 - all_1 pow 2",
		"( all_0i - all_1i ) * ( all_1i + all_0i ) = all_1 pow 2 - all_0 pow 2",

		"( all_0i - all_1i ) * ( all_0i - all_1i ) = all_0 pow 2 + all_1 pow 2 - 2 * all_0 * all_1",
		"( all_0i - all_1i ) * ( all_1i - all_0i ) = all_0 pow 2 + all_1 pow 2 - 2 * all_0 * all_1",

		// - Four
		"( all_0i + all_1i ) * ( all_2i + all_3i ) = all_0 * all_2 + all_1 * all_2 + all_0 * all_3 + all_1 * all_3",
		"( all_0i + all_1i ) * ( all_2i - all_3i ) = all_0 * all_2 + all_1 * all_2 - all_0 * all_3 - all_1 * all_3",
		"( all_0i - all_1i ) * ( all_2i + all_3i ) = all_0 * all_2 - all_1 * all_2 + all_0 * all_3 - all_1 * all_3",
		"( all_0i - all_1i ) * ( all_2i - all_3i ) = all_0 * all_2 - all_1 * all_2 - all_0 * all_3 - all_1 * all_3",

		// - Three
		"all_0i * ( all_1i + all_2i ) = all_0 * all_1 + all_0 * all_2",
		"all_0i * ( all_1i - all_2i ) = all_0 * all_1 - all_0 * all_2",
		"( all_0i + all_1i ) * all_2i = all_2 * all_0 + all_2 * all_1",
		"( all_0i - all_1i ) * all_2i = all_2 * all_0 - all_2 * all_1",

		// Delete Rules:
		"0 + all_0 = all_0",
		"0 * all_0 = 0",
		"all_0 - all_0 = 0",
		"all_0 / 1 = all_0",

		// All vars should have x * var
		"all_0 + var_1 = all_0 + 1 * var_1",
		"var_0 + all_1 = 1 * var_0 + all_1",

		// Remove Division
		"all_0 / num_1 = 1 / num_1 * all_0",
		"all_0 / all_1 = all_0 * all_1 pow -1",

		//Remove Subtraction
		"all_0 - all_1 = all_0 + -1 * all_1",

		// Merge Var
		"num_0 * all_1 + num_2 * all_1 = ( num_0 + num_2 ) * all_1",
		"all_0 + num_1 * all_2 + num_3 * all_2 = all_0 + ( num_1 + num_3 ) * all_2",

		// Merge Pow
		"all_0 pow num_1 * all_0 = all_0 pow ( num_1 + 1 )",
		"all_0 pow num_1 pow num_1 = all_0 pow ( num_1 * 2 )",
		"all_0 pow num_1 pow num_2 = all_0 pow ( num_1 * num_2 )",

		// Merge Mul
		"all_0 * all_0 = all_0 pow 2",
		"all_0 * all_1 * all_1 = all_0 * all_1 pow 2", // Edge

		// Merge Addition
		"num_0 * all_1 + all_1 = ( num_0 + 1 ) * all_1",
		"all_0 + num_1 * all_2 + all_2 = all_0 + ( num_1 + 1 ) * all_2", // Edge
		"all_0 + all_0 = 2 * all_0",
		"all_0 + all_1 + all_1 = all_0 + 2 * all_1", // Edge
	}

	for _, ruleString := range ruleStrings {
		searchRoot, replaceRoot := parseReplaceRule(ruleString)
		simpPatterns = append(simpPatterns, replace(searchRoot, replaceRoot, ruleString))
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
	index := int64(len(*dataBuffer))
	for i, testnode := range *dataBuffer {
		if current.equalDeep(&testnode) {
			index = int64(i)
			break
		}
	}

	id, err := strconv.ParseInt(string(search.data[0]), 10, 64)
	handelError(err)
	if id >= int64(len(*dataBuffer)) {
		*dataBuffer = append(*dataBuffer, *current)
	}

	debug := id == index || ((len(search.data) == 2 && search.data[1] == 'i') && id >= index)
	return debug
}

func replace(search *node, replace *node, patternString string) simpPattern {
	return simpPattern{
		func(root *node) bool {
			return checkReplace(root, search, &replaceDataBuffer{})
		},
		func(root *node) *node {
			dataBuffer := &replaceDataBuffer{}
			setReplaceDataBuffer(root, search, dataBuffer)

			newRoot := replace.copyDeep()
			replaceNodes(newRoot, dataBuffer)
			return newRoot
		},
		patternString,
	}
}
func checkReplace(current *node, search *node, dataBuffer *replaceDataBuffer) bool {
	if len(search.childs) > len(current.childs) {
		return false
	}
	for i, searchChild := range search.childs {
		if !checkReplace(current.childs[i], searchChild, dataBuffer) {
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
	for i, searchChild := range search.childs {
		setReplaceDataBuffer(current.childs[i], searchChild, dataBuffer)
	}

	if search.hasFlag(flagRulePart) {
		dataBuffer.checkAndSet(current, search)
	}
}
func replaceNodes(current *node, dataBuffer *replaceDataBuffer) {
	for _, child := range current.childs {
		replaceNodes(child, dataBuffer)
	}

	if current.hasFlag(flagRulePart) {
		id, err := strconv.ParseInt(string(current.data[0]), 10, 32)
		handelError(err)
		*current = (*dataBuffer)[id]
	}
}

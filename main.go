package AlgebraCalculator

var initilized bool

func init() {
	if !initilized {
		InitLog()
		initTerm()

		initParser()
		initSimplifyer()
		initTermFunctions()
		initSolve()
		initVector()
		initSort()
		initReplace()
		GetLog()

		initilized = true
	}
}

type Result struct {
	AnswerStrings map[int]string
	TermStrings   map[int]string
	Log           []string
}

var terms []*term

func Calculate(termStrings ...string) Result {
	terms = nil

	result := Result{
		AnswerStrings: map[int]string{},
		TermStrings:   map[int]string{},
	}

	for i, termString := range termStrings {
		if termString == "" {
			continue
		}

		term, err := parseTerm(termString)
		result.Log = append(result.Log, GetLog())
		if handelError(err) {
			result.Log = append(result.Log, GetLog())
			continue
		}

		simplifyRoot(term.root)
		if r := recover(); r != nil {
			handelError(newError(errorTypSolving, errorCriticalLevelPartial, "Some Error accured!"))
			result.Log = append(result.Log, GetLog())
			continue
		}
		result.Log = append(result.Log, GetLog())

		term.root.print()
		result.AnswerStrings[i] = GetLog()

		term.print()
		result.TermStrings[i] = GetLog()

		terms = append(terms, term)
	}
	return result
}

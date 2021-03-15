package V4

import "AlgebraCalculator/log"

var initilized bool

func init() {
	if !initilized {
		log.InitLog()
		initTerm()

		initSimplifyer()
		initSolve()
		initVector()
		initSort()
		initReplace()
		log.GetLog()

		initilized = true
	}
}

type Result struct {
	AnswerStrings map[int]string
	TermStrings   map[int]string
	Log           []string
}

func Calculate(termStrings ...string) Result {
	result := Result{
		AnswerStrings: map[int]string{},
		TermStrings:   map[int]string{},
	}

	for i, termString := range termStrings {
		if termString == "" {
			continue
		}

		term, err := parseTerm(termString)
		result.Log = append(result.Log, log.GetLog())
		if handelError(err) {
			result.Log = append(result.Log, log.GetLog())
			continue
		}

		simplifyRoot(term.root)
		if r := recover(); r != nil {
			handelError(newError(errorTypSolving, errorCriticalLevelPartial, "Some Error accured!"))
			result.Log = append(result.Log, log.GetLog())
			continue
		}

		term.root.print()
		result.AnswerStrings[i] = log.GetLog()

		term.print()
		result.TermStrings[i] = log.GetLog()
	}
	return result
}

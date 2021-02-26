package V2

const (
	RankAppend        = 1
	RankAddSub        = 2
	RankMul           = 3
	RankPow           = 4
	RankFunc          = 5
	RankSubOpperation = 6
	RankTerm          = 7
	RankMax           = 7
)

type SolvableTermPart interface {
	TermPart
	solve(term *Term, index int) bool
	getRank() int
}

func solveTerm(term Term) Term {

	// Barces
	braceLevel := 0
	var openBraces []int
	var closeBraces []int
	for i, termPart := range term.parts {
		if termPart.getType() == TypBrace {

			if termPart.(Brace).opening {
				if braceLevel == 0 {
					openBraces = append(openBraces, i)
				}
				braceLevel++
			} else {
				braceLevel--
				if braceLevel == 0 {
					closeBraces = append(closeBraces, i)
				}
			}
		}
	}

	braceLengthSum := 0
	for i, openIndex := range openBraces {
		closeIndex := closeBraces[i]

		openIndex -= braceLengthSum
		closeIndex -= braceLengthSum

		subTerm := solveTerm(term.getSub(openIndex+1, closeIndex-1))
		term.setSub(openIndex, closeIndex, subTerm)

		braceLengthSum += closeIndex - openIndex
	}

	var executionOrder []int
	for i := RankMax; i >= 1; i-- {
		for j, termPart := range term.parts {

			if termPart.isSolvable() && termPart.(SolvableTermPart).getRank() == i {
				executionOrder = append(executionOrder, j)
			}
		}
	}

	reRun := false
	for _, index := range executionOrder {
		shouldReRun := term.parts[index].(SolvableTermPart).solve(&term, index)
		if shouldReRun {
			reRun = true
		}
	}

	if reRun {
		term = solveTerm(term)
	}

	return term
}

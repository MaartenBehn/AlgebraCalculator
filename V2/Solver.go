package V2

const (
	RankAddSub        = 1
	RankMul           = 2
	RankFunc          = 3
	RankSubOpperation = 4
	RankMax           = 4
)

type SolvableTermPart interface {
	TermPart
	solve(term *Term, index int)
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

		subTerm := solveTerm(term.getSub(openIndex, closeIndex))
		term.setSub(openIndex, closeIndex, subTerm)

		braceLengthSum += closeIndex - openIndex
	}

	return term
}

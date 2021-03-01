package V2

type Term struct {
	indexes []*int
	parts   []ITermPart
}

func NewTerm(parts []ITermPart) Term {
	term := Term{
		parts: parts,
	}
	term.updateIndexes()

	return term
}
func (t *Term) updateIndexes() {
	if len(t.indexes) != len(t.parts) {
		t.indexes = make([]*int, len(t.parts))
		for i := range t.indexes {
			index := i
			t.indexes[i] = &index
		}
	} else {
		for i := range t.indexes {
			*t.indexes[i] = i
		}
	}
}

func (t *Term) setSub(start int, end int, subTerm Term) {
	termVar := Term{}
	for i := 0; i < len(t.parts); i++ {
		termVar.parts = append(termVar.parts, t.parts[i])
		termVar.indexes = append(termVar.indexes, t.indexes[i])
	}

	newIndexes := append(termVar.indexes[:start], subTerm.indexes...)
	newParts := append(termVar.parts[:start], subTerm.parts...)

	termVar = Term{}
	for i := 0; i < len(t.parts); i++ {
		termVar.parts = append(termVar.parts, t.parts[i])
		termVar.indexes = append(termVar.indexes, t.indexes[i])
	}

	t.indexes = append(newIndexes, termVar.indexes[end+1:]...)
	t.parts = append(newParts, termVar.parts[end+1:]...)

	t.updateIndexes()
}
func (t *Term) getSub(start int, end int) Term {
	term := Term{
		indexes: t.indexes[start : end+1],
		parts:   t.parts[start : end+1],
	}
	term.updateIndexes()
	return term
}

const (
	TypVector       = 1
	TypVariable     = 2
	TypOpperator    = 3
	TypFunction     = 4
	TypBrace        = 5
	TypSubOperation = 6
	TypTermVariable = 7
)

type ITermPart interface {
	getType() int
	isSolvable() bool
	print()
	getSimplify() float64
}

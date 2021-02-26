package V2

type SubOperation struct {
	configuration string
}

func (o SubOperation) getType() int {
	return TypSubOperation
}

func (o SubOperation) getRank() int {
	return RankSubOpperation
}

func (o SubOperation) solve(term *Term, index int) {}

func (o SubOperation) print() {
	print("." + o.configuration)
}

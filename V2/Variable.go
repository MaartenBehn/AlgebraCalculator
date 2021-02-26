package V2

type Variable struct {
	name string
}

func (v Variable) getName() string {
	return v.name
}

func (v Variable) getType() int {
	return TypVariable
}

func (v Variable) print() {
	print(v.name)
}

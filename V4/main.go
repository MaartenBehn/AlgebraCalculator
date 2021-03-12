package V4

var initilized bool

func Init() {
	if !initilized {
		initTerm()
		initSolve()
		initReplace()
		initilized = true
	}
}

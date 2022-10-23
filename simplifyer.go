package AlgebraCalculator

func initSimplifyer() {
	simpPatterns = append(simpPatterns,
		simpPattern{
			func(root *node) bool {
				return root.hasFlag(flagNone) && len(root.childs) == 1
			},
			func(root *node) *node {
				return root.childs[0]
			},
			"remove none nodes",
		},
	)
}

var simpPatterns []simpPattern

type simpPattern struct {
	pattern       func(root *node) bool
	apply         func(root *node) *node
	patternString string
}

func simplifyRoot(root *node) {
	found := true
	for found {

		found = false
		for _, pattern := range simpPatterns {
			match := pattern.trySimpPattern(root)
			if match {
				root.printTree(0)
				Print("\n")
				found = true
				break
			}
		}
	}
}
func (p *simpPattern) trySimpPattern(node *node) bool {
	if p.pattern(node) {
		Println("\nPattern: " + p.patternString)
		newNode := p.apply(node)
		*node = *newNode
		return true
	}
	for _, child := range node.childs {
		if p.trySimpPattern(child) {
			return true
		}
	}
	return false
}

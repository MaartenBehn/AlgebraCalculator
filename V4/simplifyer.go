package V4

import "AlgebraCalculator/log"

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
				log.Print("\n")
				log.PrintLog()
				found = true
				break
			}
		}
	}
}
func (p *simpPattern) trySimpPattern(node *node) bool {
	if p.pattern(node) {
		log.Println("Pattern: " + p.patternString)
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

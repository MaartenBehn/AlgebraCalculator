package V4

var simpPatterns []simpPattern

type simpPattern struct {
	pattern func(root *node) bool
	apply   func(root *node) *node
}

func simplifyRoot(root *node) {
	found := true
	for found {

		found = false
		for _, pattern := range simpPatterns {
			match := pattern.trySimpPattern(root)
			if match {
				found = true
			}
		}

	}
}
func (p *simpPattern) trySimpPattern(node *node) bool {
	if p.pattern(node) {
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

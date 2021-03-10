package V4

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
func (p *simpPattern) trySimpPattern(root *node) bool {
	match := p.matchtSimpPattern(root)
	if match == nil {
		return false
	}
	newRoot := p.apply(match)
	*root = *newRoot
	return true
}
func (p *simpPattern) matchtSimpPattern(node *node) *node {
	if p.pattern(node) {
		return node
	}
	for _, child := range node.childs {
		match := p.matchtSimpPattern(child)
		if match != nil {
			return match
		}
	}
	return nil
}

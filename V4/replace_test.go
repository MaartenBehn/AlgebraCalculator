package V4

import "testing"

var testTerms = []struct {
	term string
	test func(root *node) bool
}{
	{
		"a x y = ( x + y ) * ( x + y )",
		func(root *node) bool {
			return root.data == "+" && root.childs[0].data == "+" && root.childs[0].childs[0].data == "pow"
		},
	},
	{
		"a x y = ( x * y + x * y ) * ( x * y + x * y )",
		func(root *node) bool {
			return root.data == "*" && root.childs[1].data == "pow"
		},
	},
	{
		"a t = t + t",
		func(root *node) bool {
			return root.data == "*" && root.childs[0].dataNumber == 2 && root.childs[1].data == "t"
		},
	},
	{
		"a x y = ( x + x ) * ( x + x )",
		func(root *node) bool {
			return root.data == "*" && root.childs[1].data == "pow"
		},
	},
}

func TestReplace(t *testing.T) {
	Init()

	for _, testTerm := range testTerms {
		term, err := parseTerm(testTerm.term)
		if err != nil {
			t.Error(err)
		}
		simplifyRoot(term.root)

		if !testTerm.test(term.root) {
			t.Error("Fail")
		}
	}

}

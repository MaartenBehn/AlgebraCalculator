package V4

import "testing"

func TestAddParsedNode(t *testing.T) {

	root := NewParserNode(0, 1, 1, NewNode("root"))
	current := root

	node1 := NewParserNode(5, 2, 2, NewNode("1", nodeFlagOpperator))
	addParsedNode(node1, &root, &current)
	if root.data != node1.data {
		t.Errorf("Failed")
	}

	node2 := NewParserNode(8, 2, 2, NewNode("2", nodeFlagOpperator))
	addParsedNode(node2, &root, &current)
	if root.parserChilds[0] != node2 {
		t.Errorf("Failed")
	}

	node3 := NewParserNode(4, 2, 2, NewNode("3", nodeFlagOpperator))
	addParsedNode(node3, &root, &current)
	if root.node != node3.node {
		t.Errorf("Failed")
	}
}

func TestParseRoot(t *testing.T) {

	root, _, err := parseRoot("4", "+", "4")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" {
		t.Error("Failed")
	}

	root, _, err = parseRoot("4", "+", "2", "+", "8")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[0].data != "+" {
		t.Error("Failed")
	}

	root, _, err = parseRoot("4", "+", "2", "*", "8")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[1].data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot("4", "+", "2", "*", "8", "*", "7")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[1].data != "*" || root.childs[1].childs[0].data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot("4", "+", "2", "*", "8", "*", "7", "+", "9")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[0].data != "+" || root.childs[0].childs[1].data != "*" ||
		root.childs[0].childs[1].childs[0].data != "*" {
		t.Error("Failed")
	}
}

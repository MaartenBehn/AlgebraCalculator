package V4

import "testing"

func TestAddParsedNode(t *testing.T) {

	root := newParserNode(0, 1, 1, newNode("root", 0))
	current := root

	node1 := newParserNode(5, 2, 2, newNode("1", 1, flagOperator2))
	addParsedNode(node1, &root, &current)
	if root.data != node1.data {
		t.Errorf("Failed")
	}

	node2 := newParserNode(8, 2, 2, newNode("2", 2, flagOperator2))
	addParsedNode(node2, &root, &current)
	if root.parserChilds[0] != node2 {
		t.Errorf("Failed")
	}

	node3 := newParserNode(4, 2, 2, newNode("3", 3, flagOperator2))
	addParsedNode(node3, &root, &current)
	if root.node != node3.node {
		t.Errorf("Failed")
	}
}

func TestParseRoot(t *testing.T) {
	root, _, err := parseRoot(parseTermFuncs, "4", "+", "4")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "+", "2", "+", "8")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[0].data != "+" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "+", "2", "*", "8")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[1].data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "+", "2", "*", "8", "*", "7")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[1].data != "*" || root.childs[1].childs[0].data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "+", "2", "*", "8", "*", "7", "+", "9")
	if err != nil {
		t.Error(err)
	}
	if root.data != "+" || root.childs[0].data != "+" || root.childs[0].childs[1].data != "*" ||
		root.childs[0].childs[1].childs[0].data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "-", "2")
	if err != nil {
		t.Error(err)
	}
	if root.data != "-" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "*", "2")
	if err != nil {
		t.Error(err)
	}
	if root.data != "*" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "4", "/", "2")
	if err != nil {
		t.Error(err)
	}
	if root.data != "/" {
		t.Error("Failed")
	}

	root, _, err = parseRoot(parseTermFuncs, "t")
	if err != nil {
		t.Error(err)
	}

	root, _, err = parseRoot(parseTermFuncs, "(", "t", ")")
	if err != nil {
		t.Error(err)
	}
}

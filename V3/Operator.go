package V3

import (
	"fmt"
	"log"
	"math"
)

var mathOperators = []*Operator{
	NewOperator(",", RankAppend, appandVector, simpNone),
	NewOperator("+", RankAddSub, add, simpAdd),
	NewOperator("-", RankAddSub, sub, simpNone),
	NewOperator("*", RankMul, mul, simpNone),
	NewOperator("/", RankMul, div, simpNone),
	NewOperator("pow", RankPow, pow, simpNone),
	NewOperator("dot", RankFunc, dot, simpNone),
}

type Operator struct {
	*NamedNode
	solveFunction    func(*Vector, *Vector) *Vector
	simplifyFunction func(INode) INode
}

func NewOperator(name string, rank int, solveFunction func(*Vector, *Vector) *Vector, simplifyFunction func(INode) INode) *Operator {
	return &Operator{
		NamedNode:        NewNamedNode(NewNode(TypOpperator, rank, 2), name),
		solveFunction:    solveFunction,
		simplifyFunction: simplifyFunction,
	}
}

func (o *Operator) copy() INode {
	copy := NewOperator(o.name, o.rank, o.solveFunction, o.simplifyFunction)
	copy.childs = make([]INode, len(o.childs))

	for i, child := range o.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (o *Operator) solve() {
	o.Node.solve()

	if o.childs[0].getType() == TypVector && o.childs[1].getType() == TypVector {
		result := o.solveFunction(o.childs[0].(*Vector), o.childs[1].(*Vector))
		replaceNode(o, result)
		result.childs = nil
	}
}
func (o *Operator) sort() bool {
	sorted := o.Node.sort()

	child0 := o.childs[0]
	child1 := o.childs[1]

	if child0.getType() != TypOpperator && child1.getType() != TypOpperator && (o.name == "+" || o.name == "*") {

		if child0.getDeepDefiner(false) > child1.getDeepDefiner(false) {
			o.childs[1] = child0
			o.childs[0] = child1
			return true
		}
	}

	if len(child0.getChilds()) < 2 {
		return sorted
	}

	child2 := child0.getChilds()[1]
	if child0.getType() == TypOpperator && ((o.name == "+" && child0.(INamedNode).getName() == "+") ||
		(o.name == "*" && child0.(INamedNode).getName() == "*")) {

		if child2.getDeepDefiner(false) > child1.getDeepDefiner(false) {
			child2.setParent(o)
			o.childs[1] = child2

			child1.setParent(child0)
			childs := child0.getChilds()
			childs[1] = child1
			child0.setChilds(childs)
			return true
		}
	}
	return sorted
}
func (o *Operator) print() {
	fmt.Print("(")
	if len(o.childs) < 1 {
		return
	}
	o.childs[0].print()
	fmt.Printf(" %s ", o.name)
	if len(o.childs) < 2 {
		return
	}
	o.childs[1].print()
	fmt.Print(")")
}
func (o *Operator) printTree(indentation int) {
	if len(o.childs) < 1 {
		return
	}
	o.childs[0].printTree(indentation + 1)

	printIndentation(indentation)
	fmt.Printf("%s\n", o.name)
	if len(o.childs) < 2 {
		return
	}
	o.childs[1].printTree(indentation + 1)
}

func appandVector(x *Vector, y *Vector) *Vector {
	result := NewVector(nil)
	result.append(x)
	result.append(y)
	return result
}
func add(x *Vector, y *Vector) *Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 + f2
	})
}
func sub(x *Vector, y *Vector) *Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 - f2
	})
}
func mul(x *Vector, y *Vector) *Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 * f2
	})
}
func div(x *Vector, y *Vector) *Vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 / f2
	})
}
func pow(x *Vector, y *Vector) *Vector {
	return genericOpperation2VScalar(x, y, math.Pow)
}
func dot(x *Vector, y *Vector) *Vector {
	result := NewVector(nil)

	if x.len == y.len {
		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[0] += x.values[i] * y.values[i]
		}

	} else {
		log.Panicf("Invalid vector Dimentions!")
	}

	return result
}

func simpAdd(node INode) INode {
	return node
}
func simpNone(node INode) INode {
	return node
}

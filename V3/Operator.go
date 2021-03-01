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
func (o *Operator) print() {
	o.childs[0].print()
	fmt.Printf(" %s ", o.name)
	o.childs[1].print()
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

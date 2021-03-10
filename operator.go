package AlgebraCalculator

import (
	"AlgebraCalculator/log"
	"math"
)

var mathOperators = []*operator{
	newOperator(",", rankAppend, appandVector),
	newOperator("+", rankAddSub, add),
	newOperator("-", rankAddSub, sub),
	newOperator("*", rankMul, mul),
	newOperator("/", rankMul, div),
	newOperator("pow", rankPow, pow),
	newOperator("dot", rankMathFunction, dot),
	newOperator("dist", rankMathFunction, dist),
}

type operator struct {
	*namedNode
	solveFunction func(*vector, *vector) *vector
}

func newOperator(name string, rank int, solveFunction func(*vector, *vector) *vector) *operator {
	return &operator{
		namedNode:     newNamedNode(newNode(typOpperator, rank, 2), name),
		solveFunction: solveFunction,
	}
}

func (o *operator) copy() iNode {
	copy := newOperator(o.name, o.rank, o.solveFunction)
	copy.childs = make([]iNode, len(o.childs))

	for i, child := range o.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (o *operator) check() error {
	err := o.node.check()
	if err != nil {
		return err
	}

	if len(o.childs) < 2 {
		return newError(errorTypParsing, errorCriticalLevelPartial, "Opperator has less than two children.")
	}
	return nil
}
func (o *operator) solve() bool {
	o.node.solve()

	if o.childs[0].getType() == typVector && o.childs[1].getType() == typVector {
		result := o.solveFunction(o.childs[0].(*vector), o.childs[1].(*vector))
		replaceNode(o, result)
		result.childs = nil
		return true
	}
	return false
}
func (o *operator) sort() bool {
	sorted := o.node.sort()

	if len(o.childs) < 2 {
		return sorted
	}

	child0 := o.childs[0]
	child1 := o.childs[1]

	if (o.name == "+" || o.name == "*") &&
		(child0.getType() != typOpperator || child0.(iNamedNode).getName() != o.name) &&
		(child1.getType() != typOpperator || child1.(iNamedNode).getName() != o.name) {

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
	if child0.getType() == typOpperator && ((o.name == "+" && child0.(iNamedNode).getName() == "+") ||
		(o.name == "*" && child0.(iNamedNode).getName() == "*")) {

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
func (o *operator) print() {
	if o.bracketRoot {
		log.Print("( ")
	}
	if len(o.childs) < 1 {
		return
	}

	o.childs[0].print()
	log.Printf(" %s ", o.name)

	if len(o.childs) < 2 {
		return
	}
	o.childs[1].print()

	if o.bracketRoot {
		log.Print(" )")
	}
}
func (o *operator) printTree(indentation int) {
	if len(o.childs) < 1 {
		return
	}
	o.childs[0].printTree(indentation + 1)

	printIndentation(indentation)
	log.Printf("%s\n", o.name)
	if len(o.childs) < 2 {
		return
	}
	o.childs[1].printTree(indentation + 1)
}

func appandVector(x *vector, y *vector) *vector {
	result := newVector(nil)
	result.append(x)
	result.append(y)
	return result
}
func add(x *vector, y *vector) *vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 + f2
	})
}
func sub(x *vector, y *vector) *vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 - f2
	})
}
func mul(x *vector, y *vector) *vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 * f2
	})
}
func div(x *vector, y *vector) *vector {
	return genericOpperation2VScalar(x, y, func(f1 float64, f2 float64) float64 {
		return f1 / f2
	})
}
func pow(x *vector, y *vector) *vector {
	return genericOpperation2VScalar(x, y, math.Pow)
}
func dot(x *vector, y *vector) *vector {
	result := newVector(nil)

	if x.len == y.len {
		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[0] += x.values[i] * y.values[i]
		}

	} else {
		handelError(newError(errorTypSolving, errorCriticalLevelNon, "Invalid vector Dimentions!"))
	}

	return result
}
func dist(x *vector, y *vector) *vector {
	return newVector([]float64{
		math.Abs(
			magnitude([]*vector{
				sub(x, y),
			}).values[0],
		),
	})
}

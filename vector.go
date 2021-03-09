package AlgebraCalculator

import (
	"AlgebraCalculator/log"
	"fmt"
	"math"
)

type vector struct {
	*node
	values []float64
	len    int
}

func newVector(values []float64) *vector {
	return &vector{
		node:   newNode(typVector, rankNotSolvable, 0),
		values: values,
		len:    len(values),
	}
}

// vaules ein bool
func (v *vector) getDefiner(vaules bool) string {
	if vaules {
		return v.definer + v.toString()
	}
	return v.definer
}

func (v *vector) getDeepDefiner(vaules bool) string {
	var deepDefiner string
	for _, child := range v.childs {
		deepDefiner += child.getDeepDefiner(vaules)
	}
	deepDefiner += v.getDefiner(vaules)
	return deepDefiner
}

func (v *vector) copy() iNode {
	copy := newVector(v.values)
	copy.childs = make([]iNode, len(v.childs))

	for i, child := range v.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (v vector) print() {
	v.node.print()
	log.Print(v.toString())
}
func (v vector) printTree(indentation int) {
	printIndentation(indentation)
	log.Print(v.toString())
	v.node.printTree(indentation)
}

func (v *vector) append(v2 *vector) {
	v.values = append(v.values, v2.values...)
	v.len = len(v.values)
}
func (v *vector) updateLen() {
	v.len = len(v.values)
}
func (v *vector) toString() string {
	var text string
	if v.len > 1 {
		text += "( "
	}

	for i, value := range v.values {

		if value == math.Trunc(value) {
			text += fmt.Sprintf("%.0f", value)
		} else {
			text += fmt.Sprintf("%.4f", value)
		}

		if i < v.len-1 {
			text += " , "
		}
	}

	if v.len > 1 {
		text += " )"
	}
	return text
}

func genericOpperation1V(x *vector, opperation func(float64) float64) *vector {
	result := newVector(nil)
	result.values = make([]float64, x.len)
	for i := 0; i < x.len; i++ {
		result.values[i] = opperation(x.values[i])
	}
	result.updateLen()
	return result
}
func genericOpperation2VScalar(x *vector, y *vector, opperation func(float64, float64) float64) *vector {
	result := newVector(nil)

	if x.len == y.len {

		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[i] = opperation(x.values[i], y.values[i])
		}

	} else if x.len == 1 {

		result.values = make([]float64, y.len)
		for i := 0; i < y.len; i++ {
			result.values[i] = opperation(x.values[0], y.values[i])
		}

	} else if y.len == 1 {

		result.values = make([]float64, x.len)
		for i := 0; i < x.len; i++ {
			result.values[i] = opperation(x.values[i], y.values[0])
		}

	} else {
		handelError(newError(errorTypSolving, errorCriticalLevelNon, "Invalid vector Dimentions!"))
	}

	result.updateLen()
	return result
}

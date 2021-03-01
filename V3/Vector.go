package V3

import (
	"fmt"
	"log"
	"math"
)

type Vector struct {
	*Node
	values []float64
	len    int
}

func NewVector(values []float64) *Vector {
	return &Vector{
		Node:   NewNode(TypVector, RankNone, 0),
		values: values,
		len:    len(values),
	}
}

func (v *Vector) copy() INode {
	copy := NewVector(v.values)
	copy.childs = make([]INode, len(v.childs))

	for i, child := range v.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (v Vector) print() {
	v.Node.print()

	if v.len > 1 {
		fmt.Print("( ")
	}

	for i, value := range v.values {

		if value == math.Trunc(value) {
			fmt.Printf("%.0f", value)
		} else {
			fmt.Printf("%.4f", value)
		}

		if i < v.len-1 {
			fmt.Print(" , ")
		}
	}

	if v.len > 1 {
		fmt.Print(" )")
	}
}

func (v *Vector) append(v2 *Vector) {
	v.values = append(v.values, v2.values...)
	v.len = len(v.values)
}
func (v *Vector) updateLen() {
	v.len = len(v.values)
}

func genericOpperation1V(x *Vector, opperation func(float64) float64) *Vector {
	result := NewVector(nil)
	result.values = make([]float64, x.len)
	for i := 0; i < x.len; i++ {
		result.values[i] = opperation(x.values[i])
	}
	result.updateLen()
	return result
}
func genericOpperation2VScalar(x *Vector, y *Vector, opperation func(float64, float64) float64) *Vector {
	result := NewVector(nil)

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
		log.Panicf("Invalid vector Dimentions!")
	}

	result.updateLen()
	return result
}

package main

import "AlgebraCalculator"

type Funktion struct {
	name string
	term string
}

func main() {
	AlgebraCalculator.Run("input.txt")
}

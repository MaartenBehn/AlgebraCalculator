package AlgebraCalculator

import (
	"strconv"
	"strings"
	"unicode"
)

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func isNumber(text string) bool {
	is := true
	for i, char := range text {
		if !(unicode.IsDigit(char) || (i == 0 && char == '-' && len(text) > 1)) {
			is = false
		}
	}
	return is
}

func removeEmptiStrings(strings []string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}

func getVector(part string) (*vector, error) {
	number, err := strconv.ParseFloat(part, 64)
	return newVector([]float64{number}), err
}

func getInt(rune uint8) int {
	return int(rune) - '0'
}

package V4

import (
	"strings"
	"unicode"
)

func isNumber(text string) bool {
	is := true
	for i, char := range text {
		if !(unicode.IsDigit(char) || (i == 0 && char == '-' && len(text) > 1) || (char == '.' && len(text) > 1)) {
			is = false
		}
	}
	return is
}

func splitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return removeEmptiStrings(strings.FieldsFunc(s, splitter)...)
}

func removeEmptiStrings(strings ...string) []string {
	for i := len(strings) - 1; i >= 0; i-- {
		if strings[i] == "" {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}

package V4

import "unicode"

func isNumber(text string) bool {
	is := true
	for i, char := range text {
		if !(unicode.IsDigit(char) || (i == 0 && char == '-' && len(text) > 1) || (char == '.' && len(text) > 1)) {
			is = false
		}
	}
	return is
}

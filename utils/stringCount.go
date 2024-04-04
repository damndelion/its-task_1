package utils

import "unicode/utf8"

func CountSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

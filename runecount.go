package runecount

import "unicode/utf8"

func RuneCount(s string) int {
	return len([]rune(s))
}

func RuneCount2(s string) int {
	return utf8.RuneCountInString(s)
}

package morealgo

import (
	"strings"
)

func LongestPalindrome(s string) string {
	s = strings.TrimSpace(s)

	if len(s) == 1 || len(s) == 0 {
		return s
	}

	lw := 0
	rw := 0
	testStr := ""

	for lw < len(s) {
		for rw < len(s) {
			testStr = s[lw : rw+1]
			println(testStr)
			rw++
		}
		lw++
	}
	return testStr
}

func IsPalindrome(s string) string {
	s = strings.TrimSpace(s)

	if len(s) == 1 {
		return s
	}

	if len(s) == 0 {
		return ""
	}

	return reverse(s)

}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

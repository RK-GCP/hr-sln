package longestsubstring

import (
	"strings"
)

// var sstr map[string]int

func GetLonestSubString(s string) int {

	if s == "" || len(strings.TrimSpace(s)) == 0 {
		return 0
	}

	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxlen := 0
	left := 0

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		}

		if i+1-left > maxlen {
			maxlen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxlen

}

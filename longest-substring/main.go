package main

import (
	"fmt"
	"strings"
)

func main() {
	getLonestSubString("djkjaabcabc")
}

func getLonestSubString(s string) string {
	maxss := ""
	//sarry := strings.Split(s, "")
	if len(strings.TrimSpace(s)) == 0 {
		return maxss
	}

	for i := 0; i < len(s); i++ {
		p := strings.Index(maxss, "abc")

		if p == -1 {
			fmt.Println("no match found")
		}

	}
	return maxss
}

//*incomplete

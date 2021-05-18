package main

import (
	"fmt"

	addtwonumbers "github.com/RK-GCP/hr-sln/add-two-numbers"
	longestsubstring "github.com/RK-GCP/hr-sln/longest-substring"
	"github.com/RK-GCP/hr-sln/twonumbers"
)

func main() {
	twonum()
	longSubStr()
	twonumbersAdd()
}

func twonum() {
	testnumbers := make([]int, 5)
	testnumbers[0] = 3
	testnumbers[1] = 0
	testnumbers[2] = 9
	testnumbers[3] = 7
	testnumbers[4] = 6

	targetnumber := 10

	result := twonumbers.TwoSum(testnumbers, targetnumber)
	fmt.Printf("%v\n", result)

	result2 := twonumbers.TwoSum(testnumbers, targetnumber)
	fmt.Printf("%v\n", result2)
}

func longSubStr() {
	teststr := "abcabcbb"
	ls := longestsubstring.GetLonestSubString(teststr)
	fmt.Printf("%v\n", fmt.Sprintf("longest non-repeating substr for %[1]v is %[2]v", teststr, ls))
}

func twonumbersAdd() {
	ln1 := addtwonumbers.Makeln([]int{1, 2, 7})

	ln2 := addtwonumbers.Makeln([]int{5, 2, 2})

	crln := addtwonumbers.Newll(ln1, ln2)

	str := addtwonumbers.PrintLL(crln)

	fmt.Printf("%v", str)
}

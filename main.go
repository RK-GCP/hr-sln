package main

import (
	"fmt"

	addtwonumbers "github.com/RK-GCP/hr-sln/add-two-numbers"
	longestsubstring "github.com/RK-GCP/hr-sln/longest-substring"
	"github.com/RK-GCP/hr-sln/mediansortedarray"
	"github.com/RK-GCP/hr-sln/morealgo"
	"github.com/RK-GCP/hr-sln/twonumbers"
)

func main() {
	twonum()
	longSubStr()
	twonumbersAdd()
	sortedMedianArray()
	testReverse()
	//palindrome()
	getGroupAnagrams()
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
	ls := longestsubstring.GetLongestSubString(teststr)
	fmt.Printf("%v\n", fmt.Sprintf("longest non-repeating substr for %[1]v is %[2]v", teststr, ls))
}

func twonumbersAdd() {
	ln1 := addtwonumbers.Makeln([]int{1, 9, 7})

	ln2 := addtwonumbers.Makeln([]int{5, 2, 6})

	crln := addtwonumbers.Newll(ln1, ln2)

	str := addtwonumbers.PrintLL(crln)

	fmt.Printf("%v", str)
}

func sortedMedianArray() {
	nums2 := []int{1, 2, 3}
	nums1 := []int{4, 5}

	median := mediansortedarray.FindMedianSortedArrays(nums1, nums2)

	fmt.Printf("median: %v", median)
}

func testReverse() {
	i := morealgo.Reverse(-123)
	fmt.Println(i)
}

func palindrome() {
	fmt.Println(morealgo.LongestPalindrome("babad"))
	//fmt.Println(morealgo.Reverse("abba"))
}

func getGroupAnagrams() {
	test := []string{"tea", "eat", "ate", "pin", "nip"}

	validation := morealgo.GroupAnagrams(test)

	fmt.Println(validation)

}

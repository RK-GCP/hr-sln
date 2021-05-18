package plusminus

import (
	"fmt"
)

func Plusminus(nums []int32) bool {
	result := false
	n := len(nums)
	positive := 0
	negative := 0
	zeroes := 0
	for _, val := range nums {
		if val < 0 {
			negative++
		} else if val > 0 {
			positive++
		} else {
			zeroes++
		}
	}

	fmt.Printf("%.6f\n", float64(positive)/float64(n))
	fmt.Printf("%.6f\n", float64(negative)/float64(n))
	fmt.Printf("%.6f\n", float64(zeroes)/float64(n))
	result = true
	return result
}

//https://www.hackerrank.com/challenges/plus-minus/problem

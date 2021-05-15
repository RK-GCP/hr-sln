package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter series of numbers seperate by space")
	reader := bufio.NewReader(os.Stdin)

	input, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println("no input read")
	}

	inputString := strings.TrimRight(string(input), "\r\n")

	fmt.Println(inputString)

	inputArray := strings.Split(inputString, " ")

	fmt.Printf("input %v", inputArray)

	var intArray []int32

	for _, v := range inputArray {
		intItem, err := strconv.ParseInt(v, 10, 64)

		if err != nil {
			fmt.Println("error", v)
		} else {
			int32Item := int32(intItem)

			intArray = append(intArray, int32Item)

			fmt.Printf("\n#%d\n", int32Item)
		}

	}

	fmt.Println(intArray)

	plusminus(intArray)

}

func plusminus(nums []int32) bool {
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

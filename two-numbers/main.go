package main

import "fmt"

func main() {
	testnumbers := make([]int, 5)
	testnumbers[0] = 3
	testnumbers[1] = 0
	testnumbers[2] = 9
	testnumbers[3] = 7
	testnumbers[4] = 6

	targetnumber := 10

	result := twoSum(testnumbers, targetnumber)
	fmt.Printf("%v\n", result)

	result2 := twoSum(testnumbers, targetnumber)
	fmt.Printf("%v\n", result2)
}

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}

		index[b] = i
	}
	return nil
}

func two_sum(numbers []int, tgtnumber int) []int {
	index := make(map[int]int, len(numbers))

	for position, value := range numbers {
		if i, found := index[tgtnumber-value]; found {
			return []int{i, position}
		}
		index[value] = position
	}
	return nil
}

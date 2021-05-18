package twonumbers

func TwoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}

		index[b] = i
	}
	return nil
}

func Two_sum(numbers []int, tgtnumber int) []int {
	index := make(map[int]int, len(numbers))

	for position, value := range numbers {
		if i, found := index[tgtnumber-value]; found {
			return []int{i, position}
		}
		index[value] = position
	}
	return nil
}

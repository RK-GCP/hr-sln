package morealgo

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	if x == 0 {
		return x
	}

	isNegative := false

	if x < 0 {
		isNegative = true
		x = -1 * x
	}

	m := make(map[int]string)
	numArray := strings.Split(strconv.Itoa(x), "")
	sortedKeys := make([]int, len(numArray))

	for i, v := range numArray {
		m[i+1] = v
		sortedKeys[i] = i + 1
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedKeys)))

	reversed := ""
	for _, p := range sortedKeys {
		reversed += m[p]
	}

	intToReturn, err := strconv.Atoi(reversed)
	if err != nil {
		return 0
	}

	if isNegative {
		intToReturn = -1 * intToReturn
	}

	if intToReturn > math.MaxInt32 || intToReturn < math.MinInt32 {
		intToReturn = 0
	}

	return intToReturn

}

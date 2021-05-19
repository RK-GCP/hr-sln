package mediansortedarray

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64 = float64(0)

	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return result
	}

	var allnums = make([]int, totalLen)

	copy(allnums[:len(nums1)], nums1[:])
	copy(allnums[len(nums1):], nums2[:])

	sort.Ints(allnums)

	carry := totalLen % 2
	mIndex := (totalLen / 2)
	if carry > 0 {
		mIndex += carry
		result = float64(allnums[mIndex-1])
	} else {
		result = (float64(allnums[mIndex]) + float64(allnums[mIndex-1])) / 2.0
	}

	return result
}

package Subsets

import "math"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	size := int(math.Pow(float64(2), float64(len(nums))))
	result := make([][]int, size)

	result[0] = []int{}

	idx := 1
	for i := 0; i < len(nums); i++ {
		max := idx
		for j := 0; j < max; j++ {
			result[idx] = copyAndAppend(result[j], nums[i])
			idx++
		}
	}

	return result

}

func copyAndAppend(nums []int, i int) []int {
	ary := make([]int, len(nums)+1)
	copy(ary, nums)
	ary[len(nums)] = i
	return ary
}

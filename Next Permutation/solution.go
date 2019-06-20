package Next_Permutation

import "fmt"

func nextPermutation(nums []int) {
	size := len(nums)

	for i := size - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			minIdx := findMinBiggerThan(nums[i:], nums[i-1]) + i
			fmt.Println(minIdx)
			nums[i-1], nums[minIdx] = nums[minIdx], nums[i-1]

			reverse(nums[i:])
			return
		}

	}

	reverse(nums)
}

func reverse(nums []int) {
	size := len(nums)
	for i := 0; i < size/2; i++ {
		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
	}
}

func findMinBiggerThan(nums []int, target int) int {
	minVal := nums[0]
	minIdx := 0

	for i, n := range nums {
		if n <= minVal && n > target {
			minIdx = i
			minVal = n
		}
	}
	return minIdx
}

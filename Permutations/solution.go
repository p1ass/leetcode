package Permutations

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{nil}
	}
	ans := [][]int{}

	for _, n := range nums {
		rest := remove(nums, n)
		lists := permute(rest)

		for i := 0; i < len(lists); i++ {
			lists[i] = append(lists[i], n)
			fmt.Println(lists[i])
			ans = append(ans, lists[i])
		}
	}
	return ans
}

func remove(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			result = append(result, nums[i])
		}
	}
	return result
}

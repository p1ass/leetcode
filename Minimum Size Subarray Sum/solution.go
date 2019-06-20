package Minimum_Size_Subarray_Sum

func minSubArrayLen(s int, nums []int) int {
	acc := cumSum(nums)

	ans := 1050000000

	left := 0
	for right := 1; right < len(acc); right++ {
		for acc[right]-acc[left] >= s {
			ans = min(ans, right-left)
			left++
		}
	}

	if ans == 1050000000 {
		return 0
	}
	return ans
}

func cumSum(nums []int) []int {
	ac := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		ac[i+1] = ac[i] + nums[i]
	}
	return ac
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

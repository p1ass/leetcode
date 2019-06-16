package Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([][]int, len(nums)+1)
	dp[0] = []int{-105000000, -105000000, 0}

	ans := 1

	n := len(nums)
	for j := 0; j < n; j++ {
		for i := 0; i < len(nums); i++ {
			// fmt.Printf("%#v\n",dp)
			dp[i+1] = []int{0, 0, 0}
			if nums[i] > dp[i][0] {
				dp[i+1][2] = dp[i][2] + 1
				dp[i+1][1] = dp[i][0]
				dp[i+1][0] = nums[i]
			} else {
				dp[i+1][2] = dp[i][2]
				if nums[i] <= dp[i][1] {
					dp[i+1][1] = dp[i][1]
					dp[i+1][0] = dp[i][0]
				} else {
					dp[i+1][1] = dp[i][1]
					dp[i+1][0] = nums[i]
				}
			}
			ans = max(ans, dp[i+1][2])
		}
		nums = nums[1:]
	}

	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

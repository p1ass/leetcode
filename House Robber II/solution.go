package House_Robber_II

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	dp1 := [2]int{0, 0}
	dp2 := [2]int{0, nums[0]}

	for i := 1; i < size; i++ {
		dp1 = [2]int{dp1[1], max(dp1[1], dp1[0]+nums[i])}
		if i != size-1 {
			dp2 = [2]int{dp2[1], max(dp2[1], dp2[0]+nums[i])}
		}
	}

	return max(dp1[1], dp2[1])
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

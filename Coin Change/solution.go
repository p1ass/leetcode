package Coin_Change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		// fmt.Printf("%#v\n",dp)
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

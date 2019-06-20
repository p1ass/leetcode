package Unique_Paths

func uniquePaths(m int, n int) int {
	if m*n == 0 {
		return 1
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		dp[i][0] = 1
	}

	for j := 0; j < m; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}

	return dp[n-1][m-1]
}

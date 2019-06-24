package Word_Break

func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}

	for _, d := range wordDict {
		dict[d] = true
	}

	rs := []rune(s)

	dp := make([]bool, len(rs))

	for i := 0; i < len(rs); i++ {

		if _, ok := dict[string(rs[:i+1])]; ok {
			dp[i] = true
			continue
		}

		for j := 0; j < i; j++ {
			if _, ok := dict[string(rs[j+1:i+1])]; dp[j] && ok {
				dp[i] = true
				break
			}
		}
		// fmt.Printf("%#v\n",dp)
	}
	return dp[len(rs)-1]
}

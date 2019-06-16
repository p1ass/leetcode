package Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {

	r := []rune(s)
	n := len(r)
	if n == 0 {
		return 0
	}

	first := 0
	last := 0

	cnt := map[rune]int{} //indexを保存
	cnt[r[0]] = 0

	ans := 1
	for i := 1; i < n; i++ {
		if i, ok := cnt[r[i]]; ok {
			if i >= first {
				first = i + 1
			}
		}

		cnt[r[i]] = i
		last = i
		ans = max(ans, last-first+1)
	}
	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

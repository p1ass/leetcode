package Is_Subsequence

func isSubsequence(s string, t string) bool {
	cnt := map[rune][]int{}

	for i, r := range t {
		cnt[r] = append(cnt[r], i)
	}

	currentIdx := -1
	for _, r := range s {
		indices := cnt[r]
		idx := minLargerThan(indices, currentIdx)
		if idx != 105000000 {
			currentIdx = idx
		} else {
			return false
		}

	}
	return true
}

func minLargerThan(nums []int, target int) int {
	minVal := 105000000
	for _, n := range nums {
		if minVal > n && n > target {
			minVal = n
		}
	}
	return minVal
}

package Generate_Parentheses

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	caches := make([]map[string]bool, n)
	caches[0] = map[string]bool{}

	cache := map[string]bool{}
	cache["()"] = true

	for i := 1; i < n; i++ {
		caches[i] = cache
		newCache := map[string]bool{}
		for k, _ := range cache {
			newCache["("+k+")"] = true
		}

		for j := 0; j < (i+1)/2; j++ {
			for k1, _ := range caches[j+1] {
				for k2, _ := range caches[i-j] {
					newCache[k1+k2] = true
					newCache[k2+k1] = true
				}
			}
		}

		cache = newCache
	}

	ans := []string{}

	for k, _ := range cache {
		ans = append(ans, k)
	}
	return ans
}

package Combination_Sum

func combinationSum(candidates []int, target int) [][]int {
	sorted := []int{}

	for i := len(candidates) - 1; i >= 0; i-- {
		sorted = append(sorted, candidates[i])
	}

	ans := [][]int{}

	tmpSet := calcSet(sorted, 0, target, []int{}, 0)

	ans = append(ans, tmpSet...)

	return ans

}

func calcSet(candidates []int, sum int, target int, currentSet []int, idx int) [][]int {
	// fmt.Println(currentSet)
	if sum == target {
		return [][]int{currentSet}
	}

	if sum > target {
		return nil
	}

	sumSet := [][]int{}

	for i := idx; i < len(candidates); i++ {
		tmpSet := make([]int, len(currentSet))
		copy(tmpSet, currentSet)
		tmpSet = append(tmpSet, candidates[i])

		tmpAns := calcSet(candidates, sum+candidates[i], target, tmpSet, i)

		if tmpAns != nil {
			sumSet = append(sumSet, tmpAns...)
		}
	}

	return sumSet
}

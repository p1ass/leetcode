package Subarray_Sum_Equals_K

func subarraySum(nums []int, k int) int {

	ans := 0

	cnt := map[int]int{}

	cnt[0] = 1

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if v, ok := cnt[sum-k]; ok {
			ans += v
		}

		cnt[sum]++
	}

	return ans
}

package Capacity_To_Ship_Packages_Within_D_Days

func shipWithinDays(weights []int, D int) int {

	ans := 0

	small := 0
	big := 0
	for _, w := range weights {
		big += w
		small = max(small, w)
	}

	for small <= big {
		mid := (small + big) / 2

		day := 1
		sum := 0
		for _, w := range weights {
			sum += w

			if sum > mid {
				day++
				sum = w
			}
		}

		if day <= D {
			ans = mid
			big = mid - 1
		} else {
			small = mid + 1
		}
	}

	return ans

}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

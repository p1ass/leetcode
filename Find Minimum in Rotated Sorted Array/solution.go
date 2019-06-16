package Find_Minimum_in_Rotated_Sorted_Array

func findMin(nums []int) int {
	return binarySearch(nums, 0, len(nums)-1, nums[0])
}

func binarySearch(nums []int, left, right, min int) int {
	if left == right {
		return minFunc(nums[left], min)
	}

	if right-left == 1 {
		return minFunc(minFunc(nums[left], nums[right]), min)
	}

	//fmt.Println(left, right, min)

	mid := (left + right) / 2

	if nums[mid] > min {
		return binarySearch(nums, mid+1, right, min)
	}

	if nums[mid] < min {
		return binarySearch(nums, left, mid-1, nums[mid])
	}

	return nums[mid]
}

func minFunc(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

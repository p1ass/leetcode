package Search_in_Rotated_Sorted_Array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right int, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[left] > nums[mid] && (nums[left] <= target || target < nums[mid]) || nums[left] < nums[mid] && (nums[left] <= target && target < nums[mid]) {
		return binarySearch(nums, left, mid-1, target)
	} else {
		return binarySearch(nums, mid+1, right, target)
	}
}

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        return self.binarySearch(nums, 0, len(nums) - 1, target)

    def binarySearch(self, nums, left, right, target):
        if left == right:
            if nums[left] == target:
                return left

            elif nums[left] < target:
                return left + 1

            else:
                return left

        mid = (left + right) // 2

        if nums[mid] == target:
            return mid

        if nums[mid] > target:
            return self.binarySearch(nums, left, mid, target)

        elif nums[mid] < target:
            return self.binarySearch(nums, mid + 1, right, target)

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        notZero = []

        for n in nums:
            if n != 0:
                notZero.append(n)

        for i in range(len(nums)):
            if i < len(notZero):
                nums[i] = notZero[i]

            else:
                nums[i] = 0
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ans = nums[0]
        dp = [0]
        for n in nums:
            if dp[-1] < n and dp[-1] <0:
                dp.append(n)
            else:
                dp.append(dp[-1]+n)

            ans = max(ans,dp[-1])

        return ans
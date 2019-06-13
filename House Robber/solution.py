class Solution:
    def rob(self, nums: List[int]) -> int:
        dp = [0 for i in range(len(nums) + 1)]

        ans = 0

        for i, n in enumerate(nums):
            print(dp)
            if i >= 2:
                dp[i + 1] = max(dp[i - 1] + n, dp[i - 2] + n)
            elif i == 1:
                dp[i + 1] = dp[i - 1] + n

            elif i == 0:
                dp[i + 1] = n

            ans = max(ans, dp[i + 1])
        return ans
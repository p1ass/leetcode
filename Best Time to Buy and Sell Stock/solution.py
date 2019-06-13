class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        dp = [[0, 0] for i in range(len(prices))]

        ans = 0

        for i, p in enumerate(prices):
            # print(dp)
            if i == 0:
                dp[i] = [p, p]
                continue

            if p >= dp[i - 1][0]:
                dp[i][0] = p
                dp[i][1] = dp[i - 1][1]


            elif p < dp[i - 1][1]:
                dp[i][0] = p
                dp[i][1] = p

            else:
                dp[i][0] = dp[i - 1][0]
                dp[i][1] = dp[i - 1][1]

            ans = max(ans, dp[i][0] - dp[i][1])
        return ans
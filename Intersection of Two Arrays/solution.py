class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        cnt = {}

        for n in nums1:
            cnt[n] = True

        ans = []

        for n in nums2:
            if n in cnt and n not in ans:
                ans.append(n)

        return ans
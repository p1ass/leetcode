class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        cnt = {}
        for idx, val in enumerate(nums):
            cnt[val] = idx

        print(cnt)
        for idx, val in enumerate(nums):
            if target-val in cnt and idx != cnt[target-val]:
                return [idx,cnt[target-val]]
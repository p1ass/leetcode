# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:
        return self.merge(nums)

    def merge(self,nums: List[int]) -> TreeNode:
        if len(nums) == 0:
            return None

        mid = len(nums) //2

        node = TreeNode(nums[mid])
        node.left = self.merge(nums[:mid])
        node.right = self.merge(nums[mid+1:])

        return node
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if root == None:
            return 0

        return self.dfs(root)

    def dfs(self,node:TreeNode) -> int:
        if node == None:
            return 0

        left = self.dfs(node.left)
        right = self.dfs(no1de.right)
        print(left,right)
        return max(left,right) +1
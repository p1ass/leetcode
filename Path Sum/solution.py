# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False

        return self.dfs(root,sum,0)

    def dfs(self,node : TreeNode,sum:int,currentSum:int)-> bool:
        print(currentSum)
        currentSum += node.val

        if not node.left and not node.right:
            if currentSum == sum:
                return True
            else:
                return False

        if node.left:
            left = self.dfs(node.left,sum,currentSum)
            if left:
                return True

        if node.right:
            right = self.dfs(node.right,sum,currentSum)
            if right:
                return True

        return False
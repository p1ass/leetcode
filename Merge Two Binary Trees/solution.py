# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left :TreeNode = None
        self.right : TreeNode = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if t1 and t2:
            t1.val += t2.val
        elif t2:
            t1 = t2

        self.dfs(t1,t2)
        return t1

    def dfs(self,t1:TreeNode,t2:TreeNode):
        if not t1 and not t2:
            return

        if not t2:
            return

        if not t1:
            t1 = t2
            return

        if t1.left and t2.left:
            t1.left.val += t2.left.val
            self.dfs(t1.left, t2.left)

        elif not t1.left and t2.left:
            t1.left = t2.left
            self.dfs(None, t2.left)

        if t1.right and t2.right:
            t1.right.val += t2.right.val
            self.dfs(t1.right, t2.right)

        elif not t1.right and t2.right:
            t1.right = t2.right
            self.dfs(None, t2.right)



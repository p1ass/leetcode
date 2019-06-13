# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def minDepth(self, root: TreeNode) -> int:
        queue = []
        if root == None:
            return 0

        queue.append(root)

        depth = 1
        while(True):
            n = len(queue)

            for i in range(n):
                node = queue.pop(0)

                if node.left == None and node.right == None:
                    return depth

                if node.left != None:
                    queue.append(node.left)

                if node.right != None:
                    queue.append(node.right)

            depth+=1
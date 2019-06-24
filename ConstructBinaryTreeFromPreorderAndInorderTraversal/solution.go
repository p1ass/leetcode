package ConstructBinaryTreeFromPreorderAndInorderTraversal

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	index := 0
	for i, v := range inorder {
		if v == preorder[0] {
			index = i
			break
		}
	}
	root.Val = preorder[0]
	root.Left = build(preorder[1:], inorder[:index])
	root.Right = build(preorder[index+1:], inorder[index+1:])
	return root
}

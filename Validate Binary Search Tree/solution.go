package Validate_Binary_Search_Tree

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := dfs(root.Left, -2147483649, int64(root.Val))
	right := dfs(root.Right, int64(root.Val), 2147483648)

	return left && right
}

func dfs(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	// fmt.Println(node.Val,min,max)

	val := int64(node.Val)

	if val <= min || max <= val {
		return false
	}

	left := dfs(node.Left, min, val)
	right := dfs(node.Right, val, max)

	return left && right
}

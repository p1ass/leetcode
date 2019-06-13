package mini

import "fmt"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{}
	if root == nil {
		return 0
	}

	queue = append(queue, root)

	depth := 1
	for {
		fmt.Printf("%#v\n", queue)

		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			if len(queue) >= 2 {
				queue = queue[1:]
			} else {
				queue = []*TreeNode{}
			}

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth += 1
	}
}

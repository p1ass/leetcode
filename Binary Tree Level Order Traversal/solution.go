package Binary_Tree_Level_Order_Traversal

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := &queue{}
	q.Enqueue(root)

	ans := [][]int{}

	curVals := []int{}
	for {

		if q.IsEmpty() {
			break
		}

		n := len(q.data)

		for i := 0; i < n; i++ {
			node := q.Dequeue()
			curVals = append(curVals, node.Val)

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		ans = append(ans, curVals)
		curVals = []int{}
	}

	return ans
}

type queue struct {
	data []*TreeNode
}

func (q *queue) Enqueue(val *TreeNode) {
	q.data = append(q.data, val)
}

func (q *queue) Dequeue() *TreeNode {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}

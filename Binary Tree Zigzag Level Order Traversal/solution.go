package Binary_Tree_Zigzag_Level_Order_Traversal

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	level := 0

	s := &stack{}
	nextS := &stack{}

	if root == nil {
		return ans
	}

	s.Push(root)

	currentVals := []int{}
	for !s.IsEmpty() || !nextS.IsEmpty() {
		// fmt.Printf("q: %#v\n",q.data)
		// fmt.Printf("s: %#v\n",s.data)
		if s.IsEmpty() {
			level++
			ans = append(ans, currentVals)
			currentVals = []int{}
			s, nextS = nextS, s
			continue
		}

		n := s.Pop()
		currentVals = append(currentVals, n.Val)

		if level%2 == 0 {
			if n.Left != nil {
				nextS.Push(n.Left)
			}
			if n.Right != nil {
				nextS.Push(n.Right)
			}
		} else {

			if n.Right != nil {
				nextS.Push(n.Right)
			}
			if n.Left != nil {
				nextS.Push(n.Left)
			}
		}
	}

	ans = append(ans, currentVals)

	return ans
}

type stack struct {
	data []*TreeNode
}

func (s *stack) Push(val *TreeNode) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() *TreeNode {
	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]

	return val
}

func (s *stack) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

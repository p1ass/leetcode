package Linked_List_Cycle

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head

	for {

		if slow == nil {
			return false
		} else {
			slow = slow.Next
		}

		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}

		//fmt.Println(slow, fast)
		if slow == fast {
			return true
		}
	}
}

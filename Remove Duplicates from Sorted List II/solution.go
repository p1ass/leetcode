package Remove_Duplicates_from_Sorted_List_II

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := head
	head = nil
	var tail *ListNode
	for {
		if n == nil {
			if tail != nil {
				tail.Next = nil
			}
			break
		}

		if n.Next == nil {
			if head == nil {
				head = n
				head.Next = nil
			} else {
				tail.Next = n
				tail = n
				tail.Next = nil
			}

			break
		}

		if n.Val != n.Next.Val {
			if head == nil {
				head = n
				tail = n
			} else {
				tail.Next = n
				tail = n
			}

		} else {
			for n.Next != nil && n.Val == n.Next.Val {
				n = n.Next
			}
		}

		n = n.Next

	}
	return head
}

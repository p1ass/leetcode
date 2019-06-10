package Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	seen := map[int]bool{}

	if head == nil {
		return head
	}

	n := head
	seen[n.Val] = true

	for n.Next != nil {
		if _, ok := seen[n.Next.Val]; ok {
			n.Next = n.Next.Next
			if n.Next == nil {
				break
			}
			continue
		}

		seen[n.Next.Val] = true
		n = n.Next

	}
	return head
}

package Add_Two_Numbers

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := l1
	if l1 == nil || l2 == nil {
		return nil
	}
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil && l2 != nil {
			l1.Val += l2.Val + carry
			if l1.Val >= 10 {
				l1.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
		}

		if l1.Next != nil && l2.Next == nil {
			for {
				if l1.Next == nil {
					if carry == 1 {
						l1.Next = &ListNode{1, nil}
					}
					break
				}
				l1.Next.Val += carry

				if l1.Next.Val >= 10 {
					l1.Next.Val %= 10
					carry = 1
				} else {
					carry = 0
				}
				l1 = l1.Next
			}
			break
		}

		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			for {
				if l1.Next == nil {
					if carry == 1 {
						l1.Next = &ListNode{1, nil}
					}
					break
				}

				l1.Next.Val += carry

				if l1.Next.Val >= 10 {
					l1.Next.Val %= 10
					carry = 1
				} else {
					carry = 0
				}
				l1 = l1.Next
			}
			break
		}

		if l1.Next == nil && l2.Next == nil {
			if carry == 1 {
				l1.Next = &ListNode{1, nil}
			}
			break
		}

		l1 = l1.Next
		l2 = l2.Next

	}
	return head
}

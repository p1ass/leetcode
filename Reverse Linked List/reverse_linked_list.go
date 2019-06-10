package Reverse_Linked_List

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	st := &stack{}
	n := head
	for {
		if n == nil {
			break
		}
		st.Push(n)
		n = n.Next
	}

	head = nil
	var prev *ListNode
	for !st.IsEmpty() {
		pop := st.Pop()
		pop.Next = nil

		if head == nil {
			head = pop
		} else {
			prev.Next = pop
		}
		prev = pop
	}
	return head
}

type stack struct {
	data []*ListNode
}

func (s *stack) Pop() *ListNode {
	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]
	return val
}

func (s *stack) Push(l *ListNode) {
	s.data = append(s.data, l)
}
func (s *stack) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

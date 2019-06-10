package Valid_Parentheses

type stack struct {
	data []rune
}

func (s *stack) Pop() rune {
	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]
	return val
}

func (s *stack) Push(r rune) {
	s.data = append(s.data, r)
}
func (s *stack) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	st := &stack{}
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			st.Push(r)

		default:
			if st.IsEmpty() {
				return false
			}

			pop := st.Pop()

			if r == ')' && pop == '(' || r == '}' && pop == '{' || r == ']' && pop == '[' {
				continue
			} else {
				return false
			}
		}

	}

	if st.IsEmpty() {
		return true
	}
	return false
}

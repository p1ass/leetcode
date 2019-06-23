package String_to_Integer__atoi_

import "strconv"

func myAtoi(str string) int {

	rs := []rune{}

	pn := false
	removeZero := false
	for _, r := range str {
		if r == ' ' {
			if len(rs) == 0 && !removeZero {
				continue
			} else {
				break
			}
		}

		if r == '-' {
			if pn {
				break
			}

			if len(rs) != 0 || removeZero {
				break
			}

			rs = append(rs, r)
			pn = true
			continue
		}

		if r == '+' {
			if pn {
				break
			}
			pn = true

			if len(rs) != 0 || removeZero {
				break
			}

			rs = append(rs, r)
			continue
		}

		if r == '0' && (len(rs) == 0 || len(rs) == 1 && rs[0] == '-') {
			removeZero = true
			continue
		}

		if r == '1' || r == '2' || r == '3' || r == '4' ||
			r == '5' || r == '6' || r == '7' || r == '8' || r == '9' || r == '0' {
			rs = append(rs, r)
		} else {
			if len(rs) == 0 {
				break
			}
			break
		}
	}

	if len(rs) == 0 {
		return 0
	}

	// fmt.Println(string(rs))

	// fmt.Println(rs)
	val, err := strconv.ParseInt(string(rs[:min(13, len(rs))]), 10, 64)
	if err != nil {
		return 0
	}
	// fmt.Println(val)

	if val < -2147483648 {
		return -2147483648
	}

	if val > 2147483647 {
		return 2147483647
	}

	return int(val)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

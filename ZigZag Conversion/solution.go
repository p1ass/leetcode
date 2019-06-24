package ZigZag_Conversion

import "math"

func convert(s string, numRows int) string {
	rs := []rune(s)
	size := len(rs)
	indices := []int{}

	if numRows == 1 {
		return s
	}

	for k := 0; k < numRows; k++ {
		for i := 0; i < int(math.Ceil(float64(size)/float64((numRows-1)*2))); i++ {
			if idx := k + 2*(numRows-1)*i; idx < size {
				indices = append(indices, idx)
			}
			if k != 0 && k != numRows-1 {
				bias := (numRows - 1 - k) * 2
				if idx := k + 2*(numRows-1)*i + bias; idx < size {
					indices = append(indices, idx)
				}
			}
		}
	}

	ans := []rune{}

	for _, i := range indices {
		ans = append(ans, rs[i])
	}
	return string(ans)
}

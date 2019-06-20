package K_th_Symbol_in_Grammar

import (
	"math"
)

func kthGrammar(N int, K int) int {

	K--

	swap := 0

	for i := 0; i < N-2; i++ {
		digit := int(math.Pow(2.0, float64(N-1-i)))
		if K >= digit/2 {
			K = K % (digit / 2)
			swap++
		}
	}

	if K == 0 {
		if swap%2 == 0 {
			return 0
		} else {
			return 1
		}
	}
	if K == 1 {
		if swap%2 == 0 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

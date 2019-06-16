package Pow_x__n_

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n > 0 {
		return pow(x, n, x)
	} else {
		return 1 / pow(x, -n, 1)
	}
}

func pow(x float64, n int, current float64) float64 {
	if n == 1 {
		return x
	}
	// fmt.Println(x,n,current)
	r := pow(x, n/2, current)
	if n%2 == 0 {
		return r * r
	} else {
		return r * r * x
	}
}

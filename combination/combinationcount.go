package combination

func CombinationCount(n, r int) int {
	if n < 0 {
		return 0
	}
	if r > n {
		return 0
	}
	r1 := uint(n - r)
	if r1 > uint(r) {
		r1 = uint(r)
	}
	product := uint(1)
	for i, j := uint(1), uint(n); i <= uint(r1); i, j = i+1, j-1 {
		product = product * j / i
	}
	return int(product)
}

func CombinationWithRepetitionCount(n, r int) int {
	return CombinationCount(n+r-1, r)
}

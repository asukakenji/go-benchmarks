package combination

func combination3Callback[T any](a, result []T, pos, iPosStart, iPosEnd int, f func([]T)) {
	for iPos := iPosStart; iPos < iPosEnd; iPos++ {
		result[pos] = a[iPos]
		if pos == len(result)-1 {
			f(result)
		} else {
			combination3Callback(a, result, pos+1, iPos+1, iPosEnd+1, f)
		}
	}
}

func Combination3Callback[T any](a []T, r int, f func([]T)) {
	if r < 0 || r > len(a) {
		return
	}
	if r == 0 {
		f([]T{})
		return
	}
	result := make([]T, r)
	combination3Callback(a, result, 0, 0, len(a)-r+1, f)
}

func Combination3[T any](a []T, r int) [][]T {
	count := CombinationCount(len(a), r)
	results := make([][]T, 0, count)
	Combination3Callback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

package combination

func combinationWithRepetition2Callback[T any](a, result []T, pos, iPosStart int, f func([]T)) {
	for iPos := iPosStart; iPos < len(a); iPos++ {
		result[pos] = a[iPos]
		if pos == len(result)-1 {
			f(result)
		} else {
			combinationWithRepetition2Callback(a, result, pos+1, iPos, f)
		}
	}
}

func CombinationWithRepetition2Callback[T any](a []T, r int, f func([]T)) {
	if len(a) == 0 || r < 0 {
		return
	}
	if r == 0 {
		f([]T{})
		return
	}
	result := make([]T, r)
	combinationWithRepetition2Callback(a, result, 0, 0, f)
}

func CombinationWithRepetition2[T any](a []T, r int) [][]T {
	count := CombinationWithRepetitionCount(len(a), r)
	results := make([][]T, 0, count)
	CombinationWithRepetition2Callback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

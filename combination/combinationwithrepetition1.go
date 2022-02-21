package combination

func combinationWithRepetition1Callback[T any](a, result0, result []T, f func([]T)) {
	if len(a) == 0 {
		return
	}
	if len(result) == 0 {
		f(result0)
		return
	}
	result[0] = a[0]
	combinationWithRepetition1Callback(a, result0, result[1:], f)
	combinationWithRepetition1Callback(a[1:], result0, result, f)
}

func CombinationWithRepetition1Callback[T any](a []T, r int, f func([]T)) {
	if r < 0 {
		return
	}
	result := make([]T, r)
	combinationWithRepetition1Callback(a, result, result, f)
}

func CombinationWithRepetition1[T any](a []T, r int) [][]T {
	count := CombinationWithRepetitionCount(len(a), r)
	results := make([][]T, 0, count)
	CombinationWithRepetition1Callback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

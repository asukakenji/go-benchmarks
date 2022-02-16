package combination

func combination2Callback[T any](a, result []T, rIndex int, f func([]T)) {
	if len(a) < len(result)-rIndex {
		return
	}
	if rIndex == len(result) {
		f(result)
		return
	}
	result[rIndex] = a[0]
	combination2Callback(a[1:], result, rIndex+1, f)
	combination2Callback(a[1:], result, rIndex, f)
}

func Combination2Callback[T any](a []T, r int, f func([]T)) {
	result := make([]T, r)
	combination2Callback(a, result, 0, f)
}

func Combination2[T any](a []T, r int) [][]T {
	count := CombinationCount(len(a), r)
	results := make([][]T, 0, count)
	Combination2Callback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

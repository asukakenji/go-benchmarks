package combination

func combination2aCallback[T any](result, a []T, rIndex int, f func([]T)) {
	if len(a) < len(result)-rIndex {
		return
	}
	if rIndex == len(result) {
		f(result)
		return
	}
	result[rIndex] = a[0]
	combination2aCallback(result, a[1:], rIndex+1, f)
	combination2aCallback(result, a[1:], rIndex, f)
}

func Combination2aCallback[T any](a []T, r int, f func([]T)) {
	if r < 0 {
		return
	}
	result := make([]T, r)
	combination2aCallback(result, a, 0, f)
}

func Combination2a[T any](a []T, r int) [][]T {
	count := CombinationCount(len(a), r)
	results := make([][]T, 0, count)
	Combination2aCallback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

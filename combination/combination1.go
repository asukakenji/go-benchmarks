package combination

func combination1Callback[T any](a, result0, result []T, f func([]T)) {
	if len(a) < len(result) {
		return
	}
	if len(result) == 0 {
		f(result0)
		return
	}
	result[0] = a[0]
	combination1Callback(a[1:], result0, result[1:], f)
	combination1Callback(a[1:], result0, result, f)
}

func Combination1Callback[T any](a []T, r int, f func([]T)) {
	result := make([]T, r)
	combination1Callback(a, result, result, f)
}

func Combination1[T any](a []T, r int) [][]T {
	count := CombinationCount(len(a), r)
	results := make([][]T, 0, count)
	Combination1Callback(a, r, func(x []T) {
		xCopy := make([]T, len(x))
		copy(xCopy, x)
		results = append(results, xCopy)
	})
	return results
}

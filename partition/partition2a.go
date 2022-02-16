package partition

func partition2aCallback(n, r, i int, result []int, f func([]int)) {
	n1 := n - r
	result[i] = r
	i++
	if n1 == 0 {
		f(result[:i])
		return
	}
	for r1 := 1; r1 <= n1 && r1 <= r; r1++ {
		partition2aCallback(n1, r1, i, result, f)
	}
}

func Partition2aCallback(n int, f func([]int)) {
	result := make([]int, n)
	for r := 1; r <= n; r++ {
		partition2aCallback(n, r, 0, result, f)
	}
}

func Partition2a(n int) [][]int {
	output := [][]int{}
	Partition2aCallback(n, func(x []int) {
		xCopy := make([]int, len(x))
		copy(xCopy, x)
		output = append(output, xCopy)
	})
	return output
}

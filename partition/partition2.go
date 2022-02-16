package partition

func partition2Callback(n, r, i int, result []int, f func([]int)) {
	n1 := n - r
	result[i] = r
	if n1 == 0 {
		f(result[:i+1])
		return
	}
	for r1 := 1; r1 <= n1 && r1 <= r; r1++ {
		partition2Callback(n1, r1, i+1, result, f)
	}
}

func Partition2Callback(n int, f func([]int)) {
	result := make([]int, n)
	for r := 1; r <= n; r++ {
		partition2Callback(n, r, 0, result, f)
	}
}

func Partition2(n int) [][]int {
	output := [][]int{}
	Partition2Callback(n, func(x []int) {
		xCopy := make([]int, len(x))
		copy(xCopy, x)
		output = append(output, xCopy)
	})
	return output
}

package partition

func Partition3Callback(n int, f func([]int)) {
	result := make([]int, n)

	var partition3Callback func(int, int, int)
	partition3Callback = func(n, r, i int) {
		n1 := n - r
		result[i] = r
		i++
		if n1 == 0 {
			f(result[:i])
			return
		}
		for r1 := 1; r1 <= n1 && r1 <= r; r1++ {
			partition3Callback(n1, r1, i)
		}
	}

	for r := 1; r <= n; r++ {
		partition3Callback(n, r, 0)
	}
}

func Partition3(n int) [][]int {
	output := [][]int{}
	Partition3Callback(n, func(x []int) {
		xCopy := make([]int, len(x))
		copy(xCopy, x)
		output = append(output, xCopy)
	})
	return output
}

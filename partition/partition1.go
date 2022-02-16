package partition

func partition(n, r int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	if n == r {
		return [][]int{{n}}
	}
	results := [][]int{}
	n1 := n - r
	for r1 := 1; r1 <= n1 && r1 <= r; r1++ {
		subresults := partition(n1, r1)
		for _, subresult := range subresults {
			result := []int{r}
			result = append(result, subresult...)
			results = append(results, result)
		}
	}
	return results
}

func Partition1(n int) [][]int {
	output := [][]int{}
	for r := 1; r <= n; r++ {
		results := partition(n, r)
		output = append(output, results...)
	}
	return output
}

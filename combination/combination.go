package combination

func MergeEachPairOfElementsInSetByAddition(set []uint, f func([]uint)) {
	n := len(set)
	for i := 0; i < n-1; i++ {
		set[0], set[i] = set[i], set[0]
		for j := i + 1; j < n; j++ {
			set[j] += set[0]
			f(set[1:])
			set[j] -= set[0]
		}
		set[0], set[i] = set[i], set[0]
	}
}

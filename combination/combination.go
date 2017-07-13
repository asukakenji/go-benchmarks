package combination

func MergeEachPairOfElementsInMultisetByAddition0(set []uint, f func([]uint)) {
	n := len(set)
	for i, i2 := 0, n-1; i < i2; i++ {
		set[0], set[i] = set[i], set[0]
		for j := i + 1; j < n; j++ {
			set[j] += set[0]
			f(set[1:])
			set[j] -= set[0]
		}
		set[0], set[i] = set[i], set[0]
	}
}

func MergeEachPairOfElementsInMultisetByAddition1(set []uint, f func([]uint)) {
	n := len(set)
	for i, i2 := 0, n-1; i < i2; i++ {
		t1 := set[0]
		set[0] = set[i]
		set[i] = t1
		for j := i + 1; j < n; j++ {
			set[j] += set[0]
			f(set[1:])
			set[j] -= set[0]
		}
		set[i] = set[0]
		set[0] = t1
	}
}

func MergeEachPairOfElementsInMultisetByAddition2Alt1(set []uint, f func([]uint)) {
	n := len(set)
	for i := 0; i < n-1; i++ {
		set[0], set[i] = set[i], set[0]
		for j := i + 1; j < n; j++ {
			t2 := set[j]
			set[j] += set[0]
			f(set[1:])
			set[j] = t2
		}
		set[0], set[i] = set[i], set[0]
	}
}

func MergeEachPairOfElementsInMultisetByAddition2Alt2(set []uint, f func([]uint)) {
	n := len(set)
	for i, i2 := 0, n-1; i < i2; i++ {
		set[0], set[i] = set[i], set[0]
		for j := i + 1; j < n; j++ {
			t2 := set[j]
			set[j] += set[0]
			f(set[1:])
			set[j] = t2
		}
		set[0], set[i] = set[i], set[0]
	}
}

func MergeEachPairOfElementsInMultisetByAddition3(set []uint, f func([]uint)) {
	n := len(set)
	for i, i2 := 0, n-1; i < i2; i++ {
		t1 := set[0]
		set[0] = set[i]
		set[i] = t1
		for j := i + 1; j < n; j++ {
			t2 := set[j]
			set[j] += set[0]
			f(set[1:])
			set[j] = t2
		}
		set[i] = set[0]
		set[0] = t1
	}
}

package combination_test

func CombinationWithRepetition0[T any](a []T, r int) [][]T {
	if len(a) == 0 || r < 0 {
		return [][]T{}
	}
	switch r {
	case 0:
		return CombinationWithRepetition0R0(a)
	case 1:
		return CombinationWithRepetition0R1(a)
	case 2:
		return CombinationWithRepetition0R2(a)
	case 3:
		return CombinationWithRepetition0R3(a)
	case 4:
		return CombinationWithRepetition0R4(a)
	default:
		panic("not implemented")
	}
}

func CombinationWithRepetition0R0[T any](a []T) [][]T {
	return [][]T{{}}
}

func CombinationWithRepetition0R1[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n; i1++ {
		result = append(result, []T{a[i1]})
	}
	return result
}

func CombinationWithRepetition0R2[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			result = append(result, []T{a[i1], a[i2]})
		}
	}
	return result
}

func CombinationWithRepetition0R3[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				result = append(result, []T{a[i1], a[i2], a[i3]})
			}
		}
	}
	return result
}

func CombinationWithRepetition0R4[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n; i1++ {
		for i2 := i1; i2 < n; i2++ {
			for i3 := i2; i3 < n; i3++ {
				for i4 := i3; i4 < n; i4++ {
					result = append(result, []T{a[i1], a[i2], a[i3], a[i4]})
				}
			}
		}
	}
	return result
}

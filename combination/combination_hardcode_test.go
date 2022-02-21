package combination_test

func Combination0[T any](a []T, r int) [][]T {
	if r < 0 || r > len(a) {
		return [][]T{}
	}
	if r == 0 {
		return [][]T{{}}
	}
	switch r {
	case 1:
		return Combination0R1(a)
	case 2:
		return Combination0R2(a)
	case 3:
		return Combination0R3(a)
	case 4:
		return Combination0R4(a)
	default:
		panic("not implemented")
	}
}

func Combination0R1[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n; i1++ {
		result = append(result, []T{a[i1]})
	}
	return result
}

func Combination0R2[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n-1; i1++ {
		for i2 := i1 + 1; i2 < n; i2++ {
			result = append(result, []T{a[i1], a[i2]})
		}
	}
	return result
}

func Combination0R3[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n-2; i1++ {
		for i2 := i1 + 1; i2 < n-1; i2++ {
			for i3 := i2 + 1; i3 < n; i3++ {
				result = append(result, []T{a[i1], a[i2], a[i3]})
			}
		}
	}
	return result
}

func Combination0R4[T any](a []T) [][]T {
	n := len(a)
	result := [][]T{}
	for i1 := 0; i1 < n-3; i1++ {
		for i2 := i1 + 1; i2 < n-2; i2++ {
			for i3 := i2 + 1; i3 < n-1; i3++ {
				for i4 := i3 + 1; i4 < n; i4++ {
					result = append(result, []T{a[i1], a[i2], a[i3], a[i4]})
				}
			}
		}
	}
	return result
}

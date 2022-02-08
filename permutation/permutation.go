package permutation

// permutationOrder0 is the recurive function called by
// Permutation to handle generic cases.
func permutationOrder0[T any](d int, f func([]T), s []T) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutationOrder0(d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// PermutationOrder0 applies f to each permutation of s.
// This implementation is the same as
// b1a.Permutation1ParamOrder3 (the best implementation).
func PermutationOrder0[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutationOrder0(0, f, s)
	}
}

// permutationOrder1 is the recurive function called by
// Permutation to handle generic cases.
func permutationOrder1[T any](d int, f func([]T), s []T) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutationOrder1(d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// PermutationOrder1 applies f to each permutation of s.
// This implementation is essentially the same as
// b1a.Permutation1ParamOrder3 (the best implementation),
// with a different parameter order.
func PermutationOrder1[T any](f func([]T), s []T) {
	if len(s) == 0 {
		f(s)
	} else {
		permutationOrder1(0, f, s)
	}
}

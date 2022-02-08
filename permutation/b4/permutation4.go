package b4

// permutation4ParamOrder0 is the recurive function called by
// Permutation4ParamOrder0 to handle generic cases.
func permutation4ParamOrder0[T any](s []T, d int, f func([]T)) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder0(s, d-1, f)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: sdf
func Permutation4ParamOrder0[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder0(s, len(s)-1, f)
	}
}

// permutation4ParamOrder1 is the recurive function called by
// Permutation4ParamOrder1 to handle generic cases.
func permutation4ParamOrder1[T any](s []T, f func([]T), d int) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder1(s, f, d-1)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder1 generates all permutations of s,
// and apply f to each of them.
// Order: sfd
func Permutation4ParamOrder1[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder1(s, f, len(s)-1)
	}
}

// permutation4ParamOrder2 is the recurive function called by
// Permutation4ParamOrder2 to handle generic cases.
func permutation4ParamOrder2[T any](d int, s []T, f func([]T)) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder2(d-1, s, f)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder2 generates all permutations of s,
// and apply f to each of them.
// Order: dsf
func Permutation4ParamOrder2[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder2(len(s)-1, s, f)
	}
}

// permutation4ParamOrder3 is the recurive function called by
// Permutation4ParamOrder3 to handle generic cases.
func permutation4ParamOrder3[T any](d int, f func([]T), s []T) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder3(d-1, f, s)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder3 generates all permutations of s,
// and apply f to each of them.
// Order: dfs
func Permutation4ParamOrder3[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder3(len(s)-1, f, s)
	}
}

// permutation4ParamOrder4 is the recurive function called by
// Permutation4ParamOrder4 to handle generic cases.
func permutation4ParamOrder4[T any](f func([]T), s []T, d int) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder4(f, s, d-1)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder4 generates all permutations of s,
// and apply f to each of them.
// Order: fsd
func Permutation4ParamOrder4[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder4(f, s, len(s)-1)
	}
}

// permutation4ParamOrder5 is the recurive function called by
// Permutation4ParamOrder5 to handle generic cases.
func permutation4ParamOrder5[T any](f func([]T), d int, s []T) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation4ParamOrder5(f, d-1, s)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation4ParamOrder5 generates all permutations of s,
// and apply f to each of them.
// Order: fds
func Permutation4ParamOrder5[T any](s []T, f func([]T)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation4ParamOrder5(f, len(s)-1, s)
	}
}

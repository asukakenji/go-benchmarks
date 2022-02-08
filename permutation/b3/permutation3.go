package b3

// Permutation3ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: sdf
func Permutation3ParamOrder0[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder0 func([]T, int, func([]T))
	permutation3ParamOrder0 = func(s []T, d int, f func([]T)) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder0(s, d+1, f)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder0(s, 0, f)
	}
}

// Permutation3ParamOrder1 generates all permutations of s,
// and apply f to each of them.
// Order: sfd
func Permutation3ParamOrder1[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder1 func([]T, func([]T), int)
	permutation3ParamOrder1 = func(s []T, f func([]T), d int) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder1(s, f, d+1)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder1(s, f, 0)
	}
}

// Permutation3ParamOrder2 generates all permutations of s,
// and apply f to each of them.
// Order: dsf
func Permutation3ParamOrder2[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder2 func(int, []T, func([]T))
	permutation3ParamOrder2 = func(d int, s []T, f func([]T)) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder2(d+1, s, f)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder2(0, s, f)
	}
}

// Permutation3ParamOrder3 generates all permutations of s,
// and apply f to each of them.
// Order: dfs
func Permutation3ParamOrder3[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder3 func(int, func([]T), []T)
	permutation3ParamOrder3 = func(d int, f func([]T), s []T) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder3(d+1, f, s)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder3(0, f, s)
	}
}

// Permutation3ParamOrder4 generates all permutations of s,
// and apply f to each of them.
// Order: fsd
func Permutation3ParamOrder4[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder4 func(func([]T), []T, int)
	permutation3ParamOrder4 = func(f func([]T), s []T, d int) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder4(f, s, d+1)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder4(f, s, 0)
	}
}

// Permutation3ParamOrder5 generates all permutations of s,
// and apply f to each of them.
// Order: fds
func Permutation3ParamOrder5[T any](s []T, f func([]T)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder5 func(func([]T), int, []T)
	permutation3ParamOrder5 = func(f func([]T), d int, s []T) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation3ParamOrder5(f, d+1, s)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation3ParamOrder5(f, 0, s)
	}
}

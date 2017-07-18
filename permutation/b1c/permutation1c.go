// Package b1c is different from b1 in the sense that
// closure is used, and len(s) is cached in a variable.
// This is actually slower than b1 due to the overhead of closure.
package b1c

// Permutation1ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: sdf
func Permutation1ParamOrder0(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder0 func([]uint, int, func([]uint))
	permutation1ParamOrder0 = func(s []uint, d int, f func([]uint)) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder0(s, d+1, f)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder0(s, 0, f)
	}
}

// Permutation1ParamOrder1 generates all permutations of s,
// and apply f to each of them.
// Order: sfd
func Permutation1ParamOrder1(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder1 func([]uint, func([]uint), int)
	permutation1ParamOrder1 = func(s []uint, f func([]uint), d int) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder1(s, f, d+1)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder1(s, f, 0)
	}
}

// Permutation1ParamOrder2 generates all permutations of s,
// and apply f to each of them.
// Order: dsf
func Permutation1ParamOrder2(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder2 func(int, []uint, func([]uint))
	permutation1ParamOrder2 = func(d int, s []uint, f func([]uint)) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder2(d+1, s, f)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder2(0, s, f)
	}
}

// Permutation1ParamOrder3 generates all permutations of s,
// and apply f to each of them.
// Order: dfs
func Permutation1ParamOrder3(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder3 func(int, func([]uint), []uint)
	permutation1ParamOrder3 = func(d int, f func([]uint), s []uint) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder3(d+1, f, s)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder3(0, f, s)
	}
}

// Permutation1ParamOrder4 generates all permutations of s,
// and apply f to each of them.
// Order: fsd
func Permutation1ParamOrder4(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder4 func(func([]uint), []uint, int)
	permutation1ParamOrder4 = func(f func([]uint), s []uint, d int) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder4(f, s, d+1)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder4(f, s, 0)
	}
}

// Permutation1ParamOrder5 generates all permutations of s,
// and apply f to each of them.
// Order: fds
func Permutation1ParamOrder5(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation1ParamOrder5 func(func([]uint), int, []uint)
	permutation1ParamOrder5 = func(f func([]uint), d int, s []uint) {
		if d == n1 {
			f(s)
		} else {
			for i := d; i < n; i++ {
				s[d], s[i] = s[i], s[d]
				permutation1ParamOrder5(f, d+1, s)
				s[d], s[i] = s[i], s[d]
			}
		}
	}
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder5(f, 0, s)
	}
}

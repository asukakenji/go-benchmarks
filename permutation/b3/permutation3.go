package b3

// Permutation3ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: sdf
func Permutation3ParamOrder0(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder0 func([]uint, int, func([]uint))
	permutation3ParamOrder0 = func(s []uint, d int, f func([]uint)) {
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
func Permutation3ParamOrder1(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder1 func([]uint, func([]uint), int)
	permutation3ParamOrder1 = func(s []uint, f func([]uint), d int) {
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
func Permutation3ParamOrder2(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder2 func(int, []uint, func([]uint))
	permutation3ParamOrder2 = func(d int, s []uint, f func([]uint)) {
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
func Permutation3ParamOrder3(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder3 func(int, func([]uint), []uint)
	permutation3ParamOrder3 = func(d int, f func([]uint), s []uint) {
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
func Permutation3ParamOrder4(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder4 func(func([]uint), []uint, int)
	permutation3ParamOrder4 = func(f func([]uint), s []uint, d int) {
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
func Permutation3ParamOrder5(s []uint, f func([]uint)) {
	n := len(s)
	n1 := n - 1
	var permutation3ParamOrder5 func(func([]uint), int, []uint)
	permutation3ParamOrder5 = func(f func([]uint), d int, s []uint) {
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

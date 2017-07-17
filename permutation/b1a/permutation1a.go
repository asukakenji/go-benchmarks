// Package b1a is different from b1 in the sense that
// len(s) is cached in a variable common to both branches.

package b1a

// permutation1ParamOrder0 is the recurive function called by
// Permutation1ParamOrder0 to handle generic cases.
func permutation1ParamOrder0(s []uint, d int, f func([]uint)) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder0(s, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: sdf
func Permutation1ParamOrder0(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder0(s, 0, f)
	}
}

// permutation1ParamOrder1 is the recurive function called by
// Permutation1ParamOrder1 to handle generic cases.
func permutation1ParamOrder1(s []uint, f func([]uint), d int) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder1(s, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder1 generates all permutations of s,
// and apply f to each of them.
// Order: sfd
func Permutation1ParamOrder1(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder1(s, f, 0)
	}
}

// permutation1ParamOrder2 is the recurive function called by
// Permutation1ParamOrder2 to handle generic cases.
func permutation1ParamOrder2(d int, s []uint, f func([]uint)) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder2(d+1, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder2 generates all permutations of s,
// and apply f to each of them.
// Order: dsf
func Permutation1ParamOrder2(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder2(0, s, f)
	}
}

// permutation1ParamOrder3 is the recurive function called by
// Permutation1ParamOrder3 to handle generic cases.
func permutation1ParamOrder3(d int, f func([]uint), s []uint) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder3(d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder3 generates all permutations of s,
// and apply f to each of them.
// Order: dfs
func Permutation1ParamOrder3(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder3(0, f, s)
	}
}

// permutation1ParamOrder4 is the recurive function called by
// Permutation1ParamOrder4 to handle generic cases.
func permutation1ParamOrder4(f func([]uint), s []uint, d int) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder4(f, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder4 generates all permutations of s,
// and apply f to each of them.
// Order: fsd
func Permutation1ParamOrder4(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder4(f, s, 0)
	}
}

// permutation1ParamOrder5 is the recurive function called by
// Permutation1ParamOrder5 to handle generic cases.
func permutation1ParamOrder5(f func([]uint), d int, s []uint) {
	count := len(s)
	if d == count-1 {
		f(s)
	} else {
		for i := d; i < count; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder5(f, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder5 generates all permutations of s,
// and apply f to each of them.
// Order: fds
func Permutation1ParamOrder5(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1ParamOrder5(f, 0, s)
	}
}

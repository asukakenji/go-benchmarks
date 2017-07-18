// Package b2 is different from b1 in the sense that
// the recursion is done backwards,
// and therefore the variables compare with 0 instead of len(s).
// This makes b2 about 1~2% faster than b1.
// However, the code looks much cleaner with fewer len(s) calls.
package b2

// permutation1ParamOrder0 is the recurive function called by
// Permutation1ParamOrder0 to handle generic cases.
func permutation1ParamOrder0(s []uint, d int, f func([]uint)) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder0(s, d-1, f)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder0(s, len(s)-1, f)
	}
}

// permutation1ParamOrder1 is the recurive function called by
// Permutation1ParamOrder1 to handle generic cases.
func permutation1ParamOrder1(s []uint, f func([]uint), d int) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder1(s, f, d-1)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder1(s, f, len(s)-1)
	}
}

// permutation1ParamOrder2 is the recurive function called by
// Permutation1ParamOrder2 to handle generic cases.
func permutation1ParamOrder2(d int, s []uint, f func([]uint)) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder2(d-1, s, f)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder2(len(s)-1, s, f)
	}
}

// permutation1ParamOrder3 is the recurive function called by
// Permutation1ParamOrder3 to handle generic cases.
func permutation1ParamOrder3(d int, f func([]uint), s []uint) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder3(d-1, f, s)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder3(len(s)-1, f, s)
	}
}

// permutation1ParamOrder4 is the recurive function called by
// Permutation1ParamOrder4 to handle generic cases.
func permutation1ParamOrder4(f func([]uint), s []uint, d int) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder4(f, s, d-1)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder4(f, s, len(s)-1)
	}
}

// permutation1ParamOrder5 is the recurive function called by
// Permutation1ParamOrder5 to handle generic cases.
func permutation1ParamOrder5(f func([]uint), d int, s []uint) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation1ParamOrder5(f, d-1, s)
			s[i], s[d] = s[d], s[i]
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
		permutation1ParamOrder5(f, len(s)-1, s)
	}
}

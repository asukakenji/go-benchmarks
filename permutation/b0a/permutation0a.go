// Package b0a is different from b0 in the sense that
// len(s) is cached in a variable common to both branches.

package b0a

// permutation0ParamOrder0 is the recurive function called by
// Permutation0ParamOrder0 to handle generic cases.
func permutation0ParamOrder0(s0, s []uint, f func([]uint)) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder0(s0, s[1:], f)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder0 generates all permutations of s,
// and apply f to each of them.
// Order: 0sf
func Permutation0ParamOrder0(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder0(s, s, f)
	}
}

// permutation0ParamOrder1 is the recurive function called by
// Permutation0ParamOrder1 to handle generic cases.
func permutation0ParamOrder1(s0 []uint, f func([]uint), s []uint) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder1(s0, f, s[1:])
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder1 generates all permutations of s,
// and apply f to each of them.
// Order: 0fs
func Permutation0ParamOrder1(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder1(s, f, s)
	}
}

// permutation0ParamOrder2 is the recurive function called by
// Permutation0ParamOrder2 to handle generic cases.
func permutation0ParamOrder2(s, s0 []uint, f func([]uint)) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder2(s[1:], s0, f)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder2 generates all permutations of s,
// and apply f to each of them.
// Order: s0f
func Permutation0ParamOrder2(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder2(s, s, f)
	}
}

// permutation0ParamOrder3 is the recurive function called by
// Permutation0ParamOrder3 to handle generic cases.
func permutation0ParamOrder3(s []uint, f func([]uint), s0 []uint) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder3(s[1:], f, s0)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder3 generates all permutations of s,
// and apply f to each of them.
// Order: sf0
func Permutation0ParamOrder3(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder3(s, f, s)
	}
}

// permutation0ParamOrder4 is the recurive function called by
// Permutation0ParamOrder4 to handle generic cases.
func permutation0ParamOrder4(f func([]uint), s0, s []uint) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder4(f, s0, s[1:])
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder4 generates all permutations of s,
// and apply f to each of them.
// Order: f0s
func Permutation0ParamOrder4(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder4(f, s, s)
	}
}

// permutation0ParamOrder5 is the recurive function called by
// Permutation0ParamOrder5 to handle generic cases.
func permutation0ParamOrder5(f func([]uint), s, s0 []uint) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0ParamOrder5(f, s[1:], s0)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0ParamOrder5 generates all permutations of s,
// and apply f to each of them.
// Order: fs0
func Permutation0ParamOrder5(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0ParamOrder5(f, s, s)
	}
}

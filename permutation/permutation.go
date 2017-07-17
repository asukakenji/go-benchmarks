package permutation

// permutation0 is the recurive function called by Permutation0
// to handle generic cases.
func permutation0(s0, s []uint, f func([]uint)) {
	if len(s) == 1 {
		f(s0)
	} else {
		for i, count := 0, len(s); i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0(s0, s[1:], f)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0 generates all permutations of s, and apply f to each of them.
func Permutation0(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0(s, s, f)
	}
}

// permutation0Alt1 is the recurive function called by Permutation0Alt1
// to handle generic cases.
func permutation0Alt1(s0, s []uint, f func([]uint)) {
	count := len(s)
	if count == 1 {
		f(s0)
	} else {
		for i := 0; i < count; i++ {
			s[0], s[i] = s[i], s[0]
			permutation0Alt1(s0, s[1:], f)
			s[0], s[i] = s[i], s[0]
		}
	}
}

// Permutation0Alt1 generates all permutations of s, and apply f to each of them.
// The difference between Permutation0 and Permutation0Alt1 is that in Permutation0Alt1,
// len(s) is stored in a variable common to both branches.
func Permutation0Alt1(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation0Alt1(s, s, f)
	}
}

// permutation1Alt1 is the recurive function called by Permutation1Alt1
// to handle generic cases.
func permutation1Alt1(s []uint, d int, f func([]uint)) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt1(s, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt1 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: sdf
func Permutation1Alt1(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt1(s, 0, f)
	}
}

// permutation1Alt2 is the recurive function called by Permutation1Alt2
// to handle generic cases.
func permutation1Alt2(s []uint, f func([]uint), d int) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt2(s, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt2 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: sfd
func Permutation1Alt2(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt2(s, f, 0)
	}
}

// permutation1Alt3 is the recurive function called by Permutation1Alt3
// to handle generic cases.
func permutation1Alt3(d int, s []uint, f func([]uint)) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt3(d+1, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt3 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: dsf
func Permutation1Alt3(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt3(0, s, f)
	}
}

// permutation1Alt4 is the recurive function called by Permutation1Alt4
// to handle generic cases.
func permutation1Alt4(d int, f func([]uint), s []uint) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt4(d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt4 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: dfs
func Permutation1Alt4(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt4(0, f, s)
	}
}

// permutation1Alt5 is the recurive function called by Permutation1Alt5
// to handle generic cases.
func permutation1Alt5(f func([]uint), s []uint, d int) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt5(f, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt5 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: fsd
func Permutation1Alt5(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt5(f, s, 0)
	}
}

// permutation1Alt6 is the recurive function called by Permutation1Alt6
// to handle generic cases.
func permutation1Alt6(f func([]uint), d int, s []uint) {
	if d == len(s)-1 {
		f(s)
	} else {
		for i := d; i < len(s); i++ {
			s[d], s[i] = s[i], s[d]
			permutation1Alt6(f, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1Alt6 generates all permutations of s, and apply f to each of them.
// The functions Permutation1AltN differ in parameter order.
// Order: fds
func Permutation1Alt6(s []uint, f func([]uint)) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation1Alt6(f, 0, s)
	}
}

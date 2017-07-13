package permutation

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

package permutation

// permutation is the recurive function called by
// Permutation to handle generic cases.
func permutation(f func([]uint), d int, s []uint) {
	if d == 0 {
		f(s)
	} else {
		for i := d; i >= 0; i-- {
			s[i], s[d] = s[d], s[i]
			permutation(f, d-1, s)
			s[i], s[d] = s[d], s[i]
		}
	}
}

// Permutation applies f to each permutation of s.
// This implementation is essentially the same as
// b4.Permutation4ParamOrder5 (the best implementation),
// with a different parameter order in the exported function.
func Permutation(f func([]uint), s []uint) {
	if len(s) == 0 {
		f(s)
	} else {
		permutation(f, len(s)-1, s)
	}
}

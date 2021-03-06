// Package b5 is different from previous packages in the sense that
// it implements the permutation without recursion.
// It is surprisingly slower than the recursive ones.
package b5

// Permutation5Inc generates all permutations of s,
// and apply f to each of them.
func Permutation5Inc(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
		return
	}
	a := make([]int, n)
	for d := 0; d < n; d++ {
		a[d] = d
	}
	d := 0
	// This condition is the same as:
	//     !(d == 0 && a[d] == n)
	for d != 0 || a[d] != n {
		if a[d] < n {
			if d < n-1 {
				s[a[d]], s[d] = s[d], s[a[d]]
				d++
				continue
			}
			// The following sections refers to d == n
			f(s)
			a[d]++
			continue
		}
		// The following section refers to a[d] == n:
		a[d] = d
		d--
		s[a[d]], s[d] = s[d], s[a[d]]
		a[d]++
	}
}

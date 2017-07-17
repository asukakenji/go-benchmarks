// Package b1b is different from b1 in the sense that
// n = len(s) is passed as a parameter.
// This is actually slower than b1 according to the benchmark results.
package b1b

// permutation1ParamOrder00 is the recurive function called by
// Permutation1ParamOrder00 to handle generic cases.
func permutation1ParamOrder00(n int, s []uint, d int, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder00(n, s, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder00 generates all permutations of s,
// and apply f to each of them.
// Order: nsdf
func Permutation1ParamOrder00(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder00(n, s, 0, f)
	}
}

// permutation1ParamOrder01 is the recurive function called by
// Permutation1ParamOrder01 to handle generic cases.
func permutation1ParamOrder01(s []uint, n, d int, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder01(s, n, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder01 generates all permutations of s,
// and apply f to each of them.
// Order: sndf
func Permutation1ParamOrder01(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder01(s, n, 0, f)
	}
}

// permutation1ParamOrder02 is the recurive function called by
// Permutation1ParamOrder02 to handle generic cases.
func permutation1ParamOrder02(s []uint, d, n int, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder02(s, d+1, n, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder02 generates all permutations of s,
// and apply f to each of them.
// Order: sdnf
func Permutation1ParamOrder02(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder02(s, 0, n, f)
	}
}

// permutation1ParamOrder03 is the recurive function called by
// Permutation1ParamOrder03 to handle generic cases.
func permutation1ParamOrder03(s []uint, d int, f func([]uint), n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder03(s, d+1, f, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder03 generates all permutations of s,
// and apply f to each of them.
// Order: sdfn
func Permutation1ParamOrder03(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder03(s, 0, f, n)
	}
}

// permutation1ParamOrder10 is the recurive function called by
// Permutation1ParamOrder10 to handle generic cases.
func permutation1ParamOrder10(n int, s []uint, f func([]uint), d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder10(n, s, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder10 generates all permutations of s,
// and apply f to each of them.
// Order: nsfd
func Permutation1ParamOrder10(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder10(n, s, f, 0)
	}
}

// permutation1ParamOrder11 is the recurive function called by
// Permutation1ParamOrder11 to handle generic cases.
func permutation1ParamOrder11(s []uint, n int, f func([]uint), d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder11(s, n, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder11 generates all permutations of s,
// and apply f to each of them.
// Order: snfd
func Permutation1ParamOrder11(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder11(s, n, f, 0)
	}
}

// permutation1ParamOrder12 is the recurive function called by
// Permutation1ParamOrder12 to handle generic cases.
func permutation1ParamOrder12(s []uint, f func([]uint), n, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder12(s, f, n, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder12 generates all permutations of s,
// and apply f to each of them.
// Order: sfnd
func Permutation1ParamOrder12(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder12(s, f, n, 0)
	}
}

// permutation1ParamOrder13 is the recurive function called by
// Permutation1ParamOrder13 to handle generic cases.
func permutation1ParamOrder13(s []uint, f func([]uint), d, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder13(s, f, d+1, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder13 generates all permutations of s,
// and apply f to each of them.
// Order: sfdn
func Permutation1ParamOrder13(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder13(s, f, 0, n)
	}
}

// permutation1ParamOrder20 is the recurive function called by
// Permutation1ParamOrder20 to handle generic cases.
func permutation1ParamOrder20(n, d int, s []uint, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder20(n, d+1, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder20 generates all permutations of s,
// and apply f to each of them.
// Order: ndsf
func Permutation1ParamOrder20(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder20(n, 0, s, f)
	}
}

// permutation1ParamOrder21 is the recurive function called by
// Permutation1ParamOrder21 to handle generic cases.
func permutation1ParamOrder21(d, n int, s []uint, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder21(d+1, n, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder21 generates all permutations of s,
// and apply f to each of them.
// Order: dnsf
func Permutation1ParamOrder21(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder21(0, n, s, f)
	}
}

// permutation1ParamOrder22 is the recurive function called by
// permutation1ParamOrder22 to handle generic cases.
func permutation1ParamOrder22(d int, s []uint, n int, f func([]uint)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder22(d+1, s, n, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder22 generates all permutations of s,
// and apply f to each of them.
// Order: dsnf
func Permutation1ParamOrder22(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder22(0, s, n, f)
	}
}

// permutation1ParamOrder23 is the recurive function called by
// Permutation1ParamOrder23 to handle generic cases.
func permutation1ParamOrder23(d int, s []uint, f func([]uint), n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder23(d+1, s, f, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder23 generates all permutations of s,
// and apply f to each of them.
// Order: dsfn
func Permutation1ParamOrder23(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder23(0, s, f, n)
	}
}

// permutation1ParamOrder30 is the recurive function called by
// Permutation1ParamOrder30 to handle generic cases.
func permutation1ParamOrder30(n, d int, f func([]uint), s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder30(n, d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder30 generates all permutations of s,
// and apply f to each of them.
// Order: ndfs
func Permutation1ParamOrder30(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder30(n, 0, f, s)
	}
}

// permutation1ParamOrder31 is the recurive function called by
// Permutation1ParamOrder31 to handle generic cases.
func permutation1ParamOrder31(d, n int, f func([]uint), s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder31(d+1, n, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder31 generates all permutations of s,
// and apply f to each of them.
// Order: dnfs
func Permutation1ParamOrder31(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder31(0, n, f, s)
	}
}

// permutation1ParamOrder32 is the recurive function called by
// Permutation1ParamOrder32 to handle generic cases.
func permutation1ParamOrder32(d int, f func([]uint), n int, s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder32(d+1, f, n, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder32 generates all permutations of s,
// and apply f to each of them.
// Order: dfns
func Permutation1ParamOrder32(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder32(0, f, n, s)
	}
}

// permutation1ParamOrder33 is the recurive function called by
// Permutation1ParamOrder33 to handle generic cases.
func permutation1ParamOrder33(d int, f func([]uint), s []uint, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder33(d+1, f, s, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder33 generates all permutations of s,
// and apply f to each of them.
// Order: dfsn
func Permutation1ParamOrder33(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder33(0, f, s, n)
	}
}

// permutation1ParamOrder40 is the recurive function called by
// Permutation1ParamOrder40 to handle generic cases.
func permutation1ParamOrder40(n int, f func([]uint), s []uint, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder40(n, f, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder40 generates all permutations of s,
// and apply f to each of them.
// Order: nfsd
func Permutation1ParamOrder40(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder40(n, f, s, 0)
	}
}

// permutation1ParamOrder41 is the recurive function called by
// Permutation1ParamOrder41 to handle generic cases.
func permutation1ParamOrder41(f func([]uint), n int, s []uint, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder41(f, n, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder41 generates all permutations of s,
// and apply f to each of them.
// Order: fnsd
func Permutation1ParamOrder41(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder41(f, n, s, 0)
	}
}

// permutation1ParamOrder42 is the recurive function called by
// Permutation1ParamOrder42 to handle generic cases.
func permutation1ParamOrder42(f func([]uint), s []uint, n, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder42(f, s, n, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder42 generates all permutations of s,
// and apply f to each of them.
// Order: fsnd
func Permutation1ParamOrder42(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder42(f, s, n, 0)
	}
}

// permutation1ParamOrder43 is the recurive function called by
// Permutation1ParamOrder43 to handle generic cases.
func permutation1ParamOrder43(f func([]uint), s []uint, d, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder43(f, s, d+1, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder43 generates all permutations of s,
// and apply f to each of them.
// Order: fsdn
func Permutation1ParamOrder43(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder43(f, s, 0, n)
	}
}

// permutation1ParamOrder50 is the recurive function called by
// Permutation1ParamOrder50 to handle generic cases.
func permutation1ParamOrder50(n int, f func([]uint), d int, s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder50(n, f, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder50 generates all permutations of s,
// and apply f to each of them.
// Order: nfds
func Permutation1ParamOrder50(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder50(n, f, 0, s)
	}
}

// permutation1ParamOrder51 is the recurive function called by
// Permutation1ParamOrder51 to handle generic cases.
func permutation1ParamOrder51(f func([]uint), n, d int, s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder51(f, n, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder51 generates all permutations of s,
// and apply f to each of them.
// Order: fnds
func Permutation1ParamOrder51(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder51(f, n, 0, s)
	}
}

// permutation1ParamOrder52 is the recurive function called by
// Permutation1ParamOrder52 to handle generic cases.
func permutation1ParamOrder52(f func([]uint), d, n int, s []uint) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder52(f, d+1, n, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder52 generates all permutations of s,
// and apply f to each of them.
// Order: fdns
func Permutation1ParamOrder52(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder52(f, 0, n, s)
	}
}

// permutation1ParamOrder53 is the recurive function called by
// Permutation1ParamOrder53 to handle generic cases.
func permutation1ParamOrder53(f func([]uint), d int, s []uint, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation1ParamOrder53(f, d+1, s, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation1ParamOrder53 generates all permutations of s,
// and apply f to each of them.
// Order: fdsn
func Permutation1ParamOrder53(s []uint, f func([]uint)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation1ParamOrder53(f, 0, s, n)
	}
}

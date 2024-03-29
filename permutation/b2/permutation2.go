package b2

// permutation2ParamOrder00 is the recurive function called by
// Permutation2ParamOrder00 to handle generic cases.
func permutation2ParamOrder00[T any](n int, s []T, d int, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder00(n, s, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder00 generates all permutations of s,
// and apply f to each of them.
// Order: nsdf
func Permutation2ParamOrder00[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder00(n, s, 0, f)
	}
}

// permutation2ParamOrder01 is the recurive function called by
// Permutation2ParamOrder01 to handle generic cases.
func permutation2ParamOrder01[T any](s []T, n, d int, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder01(s, n, d+1, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder01 generates all permutations of s,
// and apply f to each of them.
// Order: sndf
func Permutation2ParamOrder01[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder01(s, n, 0, f)
	}
}

// permutation2ParamOrder02 is the recurive function called by
// Permutation2ParamOrder02 to handle generic cases.
func permutation2ParamOrder02[T any](s []T, d, n int, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder02(s, d+1, n, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder02 generates all permutations of s,
// and apply f to each of them.
// Order: sdnf
func Permutation2ParamOrder02[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder02(s, 0, n, f)
	}
}

// permutation2ParamOrder03 is the recurive function called by
// Permutation2ParamOrder03 to handle generic cases.
func permutation2ParamOrder03[T any](s []T, d int, f func([]T), n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder03(s, d+1, f, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder03 generates all permutations of s,
// and apply f to each of them.
// Order: sdfn
func Permutation2ParamOrder03[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder03(s, 0, f, n)
	}
}

// permutation2ParamOrder10 is the recurive function called by
// Permutation2ParamOrder10 to handle generic cases.
func permutation2ParamOrder10[T any](n int, s []T, f func([]T), d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder10(n, s, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder10 generates all permutations of s,
// and apply f to each of them.
// Order: nsfd
func Permutation2ParamOrder10[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder10(n, s, f, 0)
	}
}

// permutation2ParamOrder11 is the recurive function called by
// Permutation2ParamOrder11 to handle generic cases.
func permutation2ParamOrder11[T any](s []T, n int, f func([]T), d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder11(s, n, f, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder11 generates all permutations of s,
// and apply f to each of them.
// Order: snfd
func Permutation2ParamOrder11[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder11(s, n, f, 0)
	}
}

// permutation2ParamOrder12 is the recurive function called by
// Permutation2ParamOrder12 to handle generic cases.
func permutation2ParamOrder12[T any](s []T, f func([]T), n, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder12(s, f, n, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder12 generates all permutations of s,
// and apply f to each of them.
// Order: sfnd
func Permutation2ParamOrder12[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder12(s, f, n, 0)
	}
}

// permutation2ParamOrder13 is the recurive function called by
// Permutation2ParamOrder13 to handle generic cases.
func permutation2ParamOrder13[T any](s []T, f func([]T), d, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder13(s, f, d+1, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder13 generates all permutations of s,
// and apply f to each of them.
// Order: sfdn
func Permutation2ParamOrder13[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder13(s, f, 0, n)
	}
}

// permutation2ParamOrder20 is the recurive function called by
// Permutation2ParamOrder20 to handle generic cases.
func permutation2ParamOrder20[T any](n, d int, s []T, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder20(n, d+1, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder20 generates all permutations of s,
// and apply f to each of them.
// Order: ndsf
func Permutation2ParamOrder20[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder20(n, 0, s, f)
	}
}

// permutation2ParamOrder21 is the recurive function called by
// Permutation2ParamOrder21 to handle generic cases.
func permutation2ParamOrder21[T any](d, n int, s []T, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder21(d+1, n, s, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder21 generates all permutations of s,
// and apply f to each of them.
// Order: dnsf
func Permutation2ParamOrder21[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder21(0, n, s, f)
	}
}

// permutation2ParamOrder22 is the recurive function called by
// permutation2ParamOrder22 to handle generic cases.
func permutation2ParamOrder22[T any](d int, s []T, n int, f func([]T)) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder22(d+1, s, n, f)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder22 generates all permutations of s,
// and apply f to each of them.
// Order: dsnf
func Permutation2ParamOrder22[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder22(0, s, n, f)
	}
}

// permutation2ParamOrder23 is the recurive function called by
// Permutation2ParamOrder23 to handle generic cases.
func permutation2ParamOrder23[T any](d int, s []T, f func([]T), n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder23(d+1, s, f, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder23 generates all permutations of s,
// and apply f to each of them.
// Order: dsfn
func Permutation2ParamOrder23[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder23(0, s, f, n)
	}
}

// permutation2ParamOrder30 is the recurive function called by
// Permutation2ParamOrder30 to handle generic cases.
func permutation2ParamOrder30[T any](n, d int, f func([]T), s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder30(n, d+1, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder30 generates all permutations of s,
// and apply f to each of them.
// Order: ndfs
func Permutation2ParamOrder30[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder30(n, 0, f, s)
	}
}

// permutation2ParamOrder31 is the recurive function called by
// Permutation2ParamOrder31 to handle generic cases.
func permutation2ParamOrder31[T any](d, n int, f func([]T), s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder31(d+1, n, f, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder31 generates all permutations of s,
// and apply f to each of them.
// Order: dnfs
func Permutation2ParamOrder31[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder31(0, n, f, s)
	}
}

// permutation2ParamOrder32 is the recurive function called by
// Permutation2ParamOrder32 to handle generic cases.
func permutation2ParamOrder32[T any](d int, f func([]T), n int, s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder32(d+1, f, n, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder32 generates all permutations of s,
// and apply f to each of them.
// Order: dfns
func Permutation2ParamOrder32[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder32(0, f, n, s)
	}
}

// permutation2ParamOrder33 is the recurive function called by
// Permutation2ParamOrder33 to handle generic cases.
func permutation2ParamOrder33[T any](d int, f func([]T), s []T, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder33(d+1, f, s, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder33 generates all permutations of s,
// and apply f to each of them.
// Order: dfsn
func Permutation2ParamOrder33[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder33(0, f, s, n)
	}
}

// permutation2ParamOrder40 is the recurive function called by
// Permutation2ParamOrder40 to handle generic cases.
func permutation2ParamOrder40[T any](n int, f func([]T), s []T, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder40(n, f, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder40 generates all permutations of s,
// and apply f to each of them.
// Order: nfsd
func Permutation2ParamOrder40[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder40(n, f, s, 0)
	}
}

// permutation2ParamOrder41 is the recurive function called by
// Permutation2ParamOrder41 to handle generic cases.
func permutation2ParamOrder41[T any](f func([]T), n int, s []T, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder41(f, n, s, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder41 generates all permutations of s,
// and apply f to each of them.
// Order: fnsd
func Permutation2ParamOrder41[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder41(f, n, s, 0)
	}
}

// permutation2ParamOrder42 is the recurive function called by
// Permutation2ParamOrder42 to handle generic cases.
func permutation2ParamOrder42[T any](f func([]T), s []T, n, d int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder42(f, s, n, d+1)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder42 generates all permutations of s,
// and apply f to each of them.
// Order: fsnd
func Permutation2ParamOrder42[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder42(f, s, n, 0)
	}
}

// permutation2ParamOrder43 is the recurive function called by
// Permutation2ParamOrder43 to handle generic cases.
func permutation2ParamOrder43[T any](f func([]T), s []T, d, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder43(f, s, d+1, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder43 generates all permutations of s,
// and apply f to each of them.
// Order: fsdn
func Permutation2ParamOrder43[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder43(f, s, 0, n)
	}
}

// permutation2ParamOrder50 is the recurive function called by
// Permutation2ParamOrder50 to handle generic cases.
func permutation2ParamOrder50[T any](n int, f func([]T), d int, s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder50(n, f, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder50 generates all permutations of s,
// and apply f to each of them.
// Order: nfds
func Permutation2ParamOrder50[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder50(n, f, 0, s)
	}
}

// permutation2ParamOrder51 is the recurive function called by
// Permutation2ParamOrder51 to handle generic cases.
func permutation2ParamOrder51[T any](f func([]T), n, d int, s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder51(f, n, d+1, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder51 generates all permutations of s,
// and apply f to each of them.
// Order: fnds
func Permutation2ParamOrder51[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder51(f, n, 0, s)
	}
}

// permutation2ParamOrder52 is the recurive function called by
// Permutation2ParamOrder52 to handle generic cases.
func permutation2ParamOrder52[T any](f func([]T), d, n int, s []T) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder52(f, d+1, n, s)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder52 generates all permutations of s,
// and apply f to each of them.
// Order: fdns
func Permutation2ParamOrder52[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder52(f, 0, n, s)
	}
}

// permutation2ParamOrder53 is the recurive function called by
// Permutation2ParamOrder53 to handle generic cases.
func permutation2ParamOrder53[T any](f func([]T), d int, s []T, n int) {
	if d == n-1 {
		f(s)
	} else {
		for i := d; i < n; i++ {
			s[d], s[i] = s[i], s[d]
			permutation2ParamOrder53(f, d+1, s, n)
			s[d], s[i] = s[i], s[d]
		}
	}
}

// Permutation2ParamOrder53 generates all permutations of s,
// and apply f to each of them.
// Order: fdsn
func Permutation2ParamOrder53[T any](s []T, f func([]T)) {
	n := len(s)
	if n == 0 {
		f(s)
	} else {
		permutation2ParamOrder53(f, 0, s, n)
	}
}

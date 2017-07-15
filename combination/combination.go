package combination

// MergeEachPairOfElementsInMultisetByAddition0 generates all combinations
// according to the following rule, and apply f to each of them:
// 1. Select 2 elements from the multiset s
// 2. Add the selected elements and produce a sum
// 3. Remove the selected elements from the multiset s
// 4. Add the sum to the multiset s
// 5. Apply f to the result multiset s'
func MergeEachPairOfElementsInMultisetByAddition0(s []uint, f func([]uint)) {
	n := len(s)
	for i, i2 := 0, n-1; i < i2; i++ {
		s[0], s[i] = s[i], s[0]
		for j := i + 1; j < n; j++ {
			s[j] += s[0]
			f(s[1:])
			s[j] -= s[0]
		}
		s[0], s[i] = s[i], s[0]
	}
}

// MergeEachPairOfElementsInMultisetByAddition1 generates all combinations
// according to the following rule, and apply f to each of them:
// 1. Select 2 elements from the multiset s
// 2. Add the selected elements and produce a sum
// 3. Remove the selected elements from the multiset s
// 4. Add the sum to the multiset s
// 5. Apply f to the result multiset s'
func MergeEachPairOfElementsInMultisetByAddition1(s []uint, f func([]uint)) {
	n := len(s)
	for i, i2 := 0, n-1; i < i2; i++ {
		t1 := s[0]  // Different: Used t1
		s[0] = s[i] // Different: Used t1
		s[i] = t1   // Different: Used t1
		for j := i + 1; j < n; j++ {
			s[j] += s[0]
			f(s[1:])
			s[j] -= s[0]
		}
		s[i] = s[0] // Different: Used t1
		s[0] = t1   // Different: Used t1
	}
}

// MergeEachPairOfElementsInMultisetByAddition2Alt1 generates all combinations
// according to the following rule, and apply f to each of them:
// 1. Select 2 elements from the multiset s
// 2. Add the selected elements and produce a sum
// 3. Remove the selected elements from the multiset s
// 4. Add the sum to the multiset s
// 5. Apply f to the result multiset s'
func MergeEachPairOfElementsInMultisetByAddition2Alt1(s []uint, f func([]uint)) {
	n := len(s)
	for i := 0; i < n-1; i++ { // Different: Not used i2
		s[0], s[i] = s[i], s[0]
		for j := i + 1; j < n; j++ {
			t2 := s[j] // Different: Used t2
			s[j] += s[0]
			f(s[1:])
			s[j] = t2 // Different: Used t2
		}
		s[0], s[i] = s[i], s[0]
	}
}

// MergeEachPairOfElementsInMultisetByAddition2Alt2 generates all combinations
// according to the following rule, and apply f to each of them:
// 1. Select 2 elements from the multiset s
// 2. Add the selected elements and produce a sum
// 3. Remove the selected elements from the multiset s
// 4. Add the sum to the multiset s
// 5. Apply f to the result multiset s'
func MergeEachPairOfElementsInMultisetByAddition2Alt2(s []uint, f func([]uint)) {
	n := len(s)
	for i, i2 := 0, n-1; i < i2; i++ {
		s[0], s[i] = s[i], s[0]
		for j := i + 1; j < n; j++ {
			t2 := s[j] // Different: Used t2
			s[j] += s[0]
			f(s[1:])
			s[j] = t2 // Different: Used t2
		}
		s[0], s[i] = s[i], s[0]
	}
}

// MergeEachPairOfElementsInMultisetByAddition3 generates all combinations
// according to the following rule, and apply f to each of them:
// 1. Select 2 elements from the multiset s
// 2. Add the selected elements and produce a sum
// 3. Remove the selected elements from the multiset s
// 4. Add the sum to the multiset s
// 5. Apply f to the result multiset s'
func MergeEachPairOfElementsInMultisetByAddition3(s []uint, f func([]uint)) {
	n := len(s)
	for i, i2 := 0, n-1; i < i2; i++ {
		t1 := s[0]  // Different: Used t1
		s[0] = s[i] // Different: Used t1
		s[i] = t1   // Different: Used t1
		for j := i + 1; j < n; j++ {
			t2 := s[j] // Different: Used t2
			s[j] += s[0]
			f(s[1:])
			s[j] = t2 // Different: Used t2
		}
		s[i] = s[0] // Different: Used t1
		s[0] = t1   // Different: Used t1
	}
}

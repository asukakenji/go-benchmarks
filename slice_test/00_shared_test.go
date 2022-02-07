package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLen(count int, length int) []int {
	var s []int
	for i := 0; i < count; i++ {
		s = make([]int, length)
	}
	return s
}

func benchmarkMakeLenCap(count int, length, capacity int) []int {
	var s []int
	for i := 0; i < count; i++ {
		s = make([]int, length, capacity)
	}
	return s
}

func benchmarkLen(count int, slice []int) int {
	var n int
	for i := 0; i < count; i++ {
		n += len(slice)
	}
	return n
}

func benchmarkCap(count int, slice []int) int {
	var m int
	for i := 0; i < count; i++ {
		m += cap(slice)
	}
	return m
}

func benchmarkIndexSequencialAccessCalibrate(count int, slice []int) int {
	var n, index int
	for i := 0; i < count; i++ {
		n += index
		index = (index + 1) % len(slice)
	}
	return n
}

func benchmarkIndexRandomAccessCalibrate(count, increment int, slice []int) int {
	var n, index int
	for i := 0; i < count; i++ {
		n += index
		index = (index + increment) % len(slice)
	}
	return n
}

func benchmarkIndexSequencialAccess(count int, slice []int) int {
	var n, index int
	for i := 0; i < count; i++ {
		n += slice[index]
		index = (index + 1) % len(slice)
	}
	return n
}

func benchmarkIndexRandomAccess(count, increment int, slice []int) int {
	var n, index int
	for i := 0; i < count; i++ {
		n += slice[index]
		index = (index + increment) % len(slice)
	}
	return n
}

func benchmarkMakeCapAndAppend1(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, 0, capacity)
		s = append(s, x)
	}
}

func benchmarkMakeCapAndAppend2(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, 0, capacity)
		s = append(s, x, x)
	}
}

func benchmarkMakeCapAndAppend4(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, 0, capacity)
		s = append(s, x, x, x, x)
	}
}

func benchmarkMakeCapAndFillRandomByAppend1WithForIndex(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, 0, capacity)
		for j := 0; j < capacity; j++ {
			x1 := gen()
			s = append(s, x1)
		}
	}
}

func benchmarkMakeCapAndFillRandomByAppend2WithForIndex(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, 0, capacity)
		for j := 0; j < capacity; j += 2 {
			x1 := gen()
			x2 := gen()
			s = append(s, x1, x2)
		}
	}
}

func benchmarkMakeCapAndFillRandomByAppend4WithForIndex(count, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, 0, capacity)
		for j := 0; j < capacity; j += 4 {
			x1 := gen()
			x2 := gen()
			x3 := gen()
			x4 := gen()
			s = append(s, x1, x2, x3, x4)
		}
	}
}

func benchmarkMakeLenAndFillRandomByAssignment1WithForIndex(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := 0; j < length; j++ {
			x1 := gen()
			s[j] = x1
		}
	}
}

func benchmarkMakeLenAndFillRandomByAssignment2WithForIndex(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := 0; j < length; j += 2 {
			x1 := gen()
			x2 := gen()
			s[j], s[j+1] = x1, x2
		}
	}
}

func benchmarkMakeLenAndFillRandomByAssignment4WithForIndex(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := 0; j < length; j += 4 {
			x1 := gen()
			x2 := gen()
			x3 := gen()
			x4 := gen()
			s[j], s[j+1], s[j+2], s[j+3] = x1, x2, x3, x4
		}
	}
}

func benchmarkMakeLenAndFillRandomByAssignmentWithForRange1(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := range s {
			x := gen()
			s[j] = x
		}
	}
}

// TODO: What's the difference?
func benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := range s {
			x := gen()
			s[j] = x
		}
	}
}

func benchmarkCloneCommon(count int, clone func([]int) []int, slice []int) []int {
	var s []int
	for i := 0; i < count; i++ {
		s = clone(slice)
	}
	return s
}

func cloneByAppendToNil(s []int) []int {
	return append([]int(nil), s...)
}

func benchmarkCloneByAppendNToNil(count, length int) {
	benchmarkCloneCommon(count, cloneByAppendToNil, numbers[:length])
}

func cloneByMakeAndAppend(s []int) []int {
	return append(make([]int, 0, len(s)), s...)
}

func benchmarkCloneByMakeCapAndAppendN(count, length int) {
	benchmarkCloneCommon(count, cloneByMakeAndAppend, numbers[:length])
}

func cloneByMakeAndCopy(s []int) []int {
	s2 := make([]int, len(s))
	copy(s2, s)
	return s2
}

func benchmarkCloneByMakeLenAndCopy(count, length int) {
	benchmarkCloneCommon(count, cloneByMakeAndCopy, numbers[:length])
}

package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLen(count, length int) {
	for i := 0; i < count; i++ {
		_ = make([]int, length)
	}
}

func benchmarkMakeLenCap(count, length, capacity int) {
	for i := 0; i < count; i++ {
		_ = make([]int, length, capacity)
	}
}

func benchmarkMakeLenCapAndLen(count, length, capacity int) {
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		_ = len(s)
	}
}

func benchmarkMakeLenCapAndCap(count, length, capacity int) {
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		_ = cap(s)
	}
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

func benchmarkMakeLenAndFillRandomByAssignmentWithForRange2(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j, _ := range s {
			x := gen()
			s[j] = x
		}
	}
}

func benchmarkCloneCommon(count int, clone func([]int) []int, slice []int) {
	for i := 0; i < count; i++ {
		_ = clone(slice)
	}
}

func benchmarkCloneByAppendNToNil(count, length int) {
	clone := func(s []int) []int {
		return append([]int(nil), s...)
	}
	benchmarkCloneCommon(count, clone, numbers[0:length])
}

func benchmarkCloneByMakeCapAndAppendN(count, capacity int) {
	clone := func(s []int) []int {
		return append(make([]int, 0, capacity), s...)
	}
	benchmarkCloneCommon(count, clone, numbers[0:capacity])
}

func benchmarkCloneByMakeLenAndCopy(count, length int) {
	clone := func(s []int) []int {
		s2 := make([]int, length)
		copy(s2, s)
		return s2
	}
	benchmarkCloneCommon(count, clone, numbers[0:length])
}

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

func benchmarkMakeLenCapAndAppend1(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, length, capacity)
		s = append(s, x)
	}
}

func benchmarkMakeLenCapAndAppend2(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, length, capacity)
		s = append(s, x, x)
	}
}

func benchmarkMakeLenCapAndAppend4(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		x := gen()
		s := make([]int, length, capacity)
		s = append(s, x, x, x, x)
	}
}

func benchmarkMakeLenCapAndFillRandomByAppend1WithForIndex(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		for j := 0; j < capacity; j++ {
			x1 := gen()
			s = append(s, x1)
		}
	}
}

func benchmarkMakeLenCapAndFillRandomByAppend2WithForIndex(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		for j := 0; j < capacity; j += 2 {
			x1 := gen()
			x2 := gen()
			s = append(s, x1, x2)
		}
	}
}

func benchmarkMakeLenCapAndFillRandomByAppend4WithForIndex(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
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

func benchmarkCloneByAppendNToNil(count, capacity int) {
	for i := 0; i < count; i++ {
		_ = append([]int(nil), numbers[0:capacity]...)
	}
}

func benchmarkCloneByMakeCapAndAppendN(count, capacity int) {
	for i := 0; i < count; i++ {
		s := make([]int, 0, capacity)
		s = append(s, numbers[0:capacity]...)
	}
}

func benchmarkCloneByMakeLenAndCopy(count, length int) {
	for i := 0; i < count; i++ {
		s := make([]int, length)
		copy(s, numbers[0:length])
	}
}

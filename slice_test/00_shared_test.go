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
			x := gen()
			s = append(s, x)
		}
	}
}

func benchmarkMakeLenAndFillRandomByAssignmentWithForIndex(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		s := make([]int, length)
		for j := 0; j < length; j++ {
			x := gen()
			s[j] = x
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

func benchmarkCloneByAppendMElementsToNilSlice(count, capacity int) {
	for i := 0; i < count; i++ {
		_ = append([]int(nil), numbers[0:capacity]...)
	}
}

func benchmarkCloneByMakeNMAndAppendMElements(count, length, capacity int) {
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		s = append(s, numbers[0:capacity]...)
	}
}

func benchmarkCloneByMakeNMAndCopy(count, length, capacity int) {
	for i := 0; i < count; i++ {
		s := make([]int, length, capacity)
		copy(s, numbers[0:capacity])
	}
}

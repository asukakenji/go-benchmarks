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
		arr := make([]int, length, capacity)
		_ = len(arr)
	}
}

func benchmarkMakeLenCapAndCap(count, length, capacity int) {
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		_ = cap(arr)
	}
}

func benchmarkMakeLenCapAndAppend1(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen())
	}
}

func benchmarkMakeLenCapAndAppend2(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		arr = append(arr, gen(), gen())
	}
}

func benchmarkMakeLenCapAndFillByAppend1WithForIndex(count, length, capacity int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		for j := 0; j < capacity; j++ {
			arr = append(arr, gen())
		}
	}
}

func benchmarkMakeLenAndFillByAssignmentWithForIndex(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length)
		for j := 0; j < length; j++ {
			arr[j] = gen()
		}
	}
}

func benchmarkMakeLenAndFillByAssignmentWithForRange1(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length)
		for j := range arr {
			arr[j] = gen()
		}
	}
}

func benchmarkMakeLenAndFillByAssignmentWithForRange2(count, length int) {
	gen := randomIntsInTheFirstPage()
	for i := 0; i < count; i++ {
		arr := make([]int, length)
		for j, _ := range arr {
			arr[j] = gen()
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
		arr := make([]int, length, capacity)
		arr = append(arr, numbers[0:capacity]...)
	}
}

func benchmarkCloneByMakeNMAndCopy(count, length, capacity int) {
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		copy(arr, numbers[0:capacity])
	}
}

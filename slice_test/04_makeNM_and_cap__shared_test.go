package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCapAndCap(count, length, capacity int) {
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		_ = cap(arr)
	}
}

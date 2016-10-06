package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCap(count, length, capacity int) {
	for i := 0; i < count; i++ {
		_ = make([]int, length, capacity)
	}
}

package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkMakeLenCapAndLen(count, length, capacity int) {
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		_ = len(arr)
	}
}

package slice_test

////////////////
// Benchmarks //
////////////////

func benchmarkCloneByMakeNMAndCopy(count, length, capacity int) {
	for i := 0; i < count; i++ {
		arr := make([]int, length, capacity)
		copy(arr, numbers[0:capacity])
	}
}

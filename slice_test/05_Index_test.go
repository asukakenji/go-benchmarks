package slice_test

import "testing"

////////////////
// Benchmarks //
////////////////

func BenchmarkIndexSinglePageSequencialAccessCalibrate(b *testing.B) {
	benchmarkIndexSequencialAccessCalibrate(b.N, numbers[:intsPerPage])
}

func BenchmarkIndexCrossPageSequencialAccessCalibrate(b *testing.B) {
	benchmarkIndexSequencialAccessCalibrate(b.N, numbers)
}

func BenchmarkIndexSinglePageSequencialAccess(b *testing.B) {
	benchmarkIndexSequencialAccess(b.N, numbers[:intsPerPage])
}

func BenchmarkIndexCrossPageSequencialAccess(b *testing.B) {
	benchmarkIndexSequencialAccess(b.N, numbers)
}

func BenchmarkIndexSinglePageRandomAccessCalibrate(b *testing.B) {
	benchmarkIndexRandomAccessCalibrate(b.N, int(indexIncrementForSinglePage), numbers[:intsPerPage])
}

func BenchmarkIndexCrossPageRandomAccessCalibrate(b *testing.B) {
	benchmarkIndexRandomAccessCalibrate(b.N, int(indexIncrementForCrossPage), numbers)
}

func BenchmarkIndexSinglePageRandomAccess(b *testing.B) {
	benchmarkIndexRandomAccess(b.N, int(indexIncrementForSinglePage), numbers[:intsPerPage])
}

func BenchmarkIndexCrossPageRandomAccess(b *testing.B) {
	benchmarkIndexRandomAccess(b.N, int(indexIncrementForCrossPage), numbers)
}

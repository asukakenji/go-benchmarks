package benchmark

import "testing"

// Runnable benchmarks a function with the signature:
//     func()
// ID: B-0-0
func Runnable(b *testing.B, runnable func()) {
	for i, count := 0, b.N; i < count; i++ {
		runnable()
	}
}

package bcommon

import "testing"

// var (
// 	index   = uint(0)
// 	numbers = [][]uint{
// 		{14556, 3742, 9070, 26821, 11803, 31083, 28105, 7365, 32544},
// 		{31549},
// 		{1296, 28906, 9632, 27000, 26974, 13051, 4223, 7056},
// 		{25443, 4555},
// 		{17178, 14913, 3075, 11320, 4865, 22683, 21602},
// 		{30115, 6416, 1653},
// 		{15827, 32039, 14348, 9896, 10925, 10605},
// 		{22174, 31173, 8201, 31479},
// 		{27431, 3402, 11322, 20749, 18733},
// 	}
// 	n = uint(len(numbers))
// )

// func Reset() {
// 	index = 0
// }

// func NextSlice() []uint {
// 	index++
// 	if index == n {
// 		index = 0
// 	}
// 	return numbers[index]
// }

func dummySliceConsumer[T any]([]T) {}

func BenchmarkPermutation[T any](b *testing.B, target func([]T, func([]T)), size int) {
	slice := make([]T, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target(slice, dummySliceConsumer[T])
	}
}

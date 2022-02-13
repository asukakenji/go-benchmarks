package bcommon

import (
	"testing"

	"github.com/asukakenji/go-benchmarks"
)

// Exported (global) variable serving as input for some
// of the benchmarks to ensure side-effect free calls
// are not optimized away.
var Input uint64

// Exported (global) variable to store function results
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var Output int

func benchmarkLeadingZeros(b *testing.B, leadingZeros func(uint) int) {
	Input = 0x03f79d71b4ca8b09
	var s int
	for i := 0; i < b.N; i++ {
		s += leadingZeros(uint(Input) >> (uint(i) % benchmarks.SizeOfUintInBits))
	}
	Output = s
}

func benchmarkLeadingZeros8(b *testing.B, leadingZeros8 func(uint8) int) {
	Input = 0x03f79d71b4ca8b09
	var s int
	for i := 0; i < b.N; i++ {
		s += leadingZeros8(uint8(Input) >> (uint(i) % 8))
	}
	Output = s
}

func benchmarkLeadingZeros16(b *testing.B, leadingZeros16 func(uint16) int) {
	Input = 0x03f79d71b4ca8b09
	var s int
	for i := 0; i < b.N; i++ {
		s += leadingZeros16(uint16(Input) >> (uint(i) % 16))
	}
	Output = s
}

func benchmarkLeadingZeros32(b *testing.B, leadingZeros32 func(uint32) int) {
	Input = 0x03f79d71b4ca8b09
	var s int
	for i := 0; i < b.N; i++ {
		s += leadingZeros32(uint32(Input) >> (uint(i) % 32))
	}
	Output = s
}

func benchmarkLeadingZeros64(b *testing.B, leadingZeros64 func(uint64) int) {
	Input = 0x03f79d71b4ca8b09
	var s int
	for i := 0; i < b.N; i++ {
		s += leadingZeros64(uint64(Input) >> (uint(i) % 64))
	}
	Output = s
}

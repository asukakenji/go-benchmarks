package math_test

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

////////////////
// Benchmarks //
////////////////

/////////////////
// Size of int //
/////////////////

const (
	maxUint              = ^uint(0)
	logSizeOfUintInBytes = maxUint>>8&1 + maxUint>>16&1 + maxUint>>32&1
	sizeOfUintInBytes    = 1 << logSizeOfUintInBytes
	sizeOfUintInBits     = sizeOfUintInBytes << 3
)

///////////////
// Page size //
///////////////

var (
	pageSizeInBytes uint
	intsPerPage     uint
)

func init() {
	pageSizeInBytes = uint(os.Getpagesize())
	intsPerPage = pageSizeInBytes / sizeOfUintInBytes
}

////////////////////////////////////////////
// Random number generator initialization //
////////////////////////////////////////////

var (
	seed int64
	src  rand.Source
	rng  *rand.Rand
)

func init() {
	seed = time.Now().UTC().UnixNano()
	src = rand.NewSource(seed)
	rng = rand.New(src)
}

///////////////
// Data size //
///////////////

const (
	B  = 1
	kB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024 * 1024
)

const (
	memoryPerSliceInBytes = 256 * MB // Configurable
)

var (
	pagesPerSlice               uint
	intsPerSlice                uint
	indexIncrementForSinglePage uint
	indexIncrementForCrossPage  uint
)

func init() {
	var n int
	pagesPerSlice = memoryPerSliceInBytes / pageSizeInBytes
	intsPerSlice = memoryPerSliceInBytes / sizeOfUintInBytes

	// If intsPerPage is 512, we would like to generate an odd integer in the
	// range [1, 511], so that it always "generates" (in mathematical sense)
	// all indexes in modulo 512.
	// Since any odd integer in [1, 511] is any odd integer in [0, 510] + 1,
	// and that is any integer in [0, 255] * 2 + 1, we have:
	n = int(intsPerPage >> 1)
	indexIncrementForSinglePage = uint(rng.Intn(n))<<1 + 1

	// If intsPerPage is 512, and pagesPerSlice is 16, then the valid range of
	// an index is [0, 512*16). We would like to generate an odd integer in the
	// range [513, 512*16-1], so that it always "cross page" and "generate" (in
	// mathematical sense) all indexes in modulo (512*16).
	// Since any odd integer in [513, 512*16-1] is any odd integer in [512, 512*16-2] + 1,
	// and that is any odd integer in [0, 512*15-2] + 512 + 1,
	// and that is any integer in [0, 256*15-1] * 2 + 512 + 1, we have:
	n = int((intsPerPage >> 1) * (pagesPerSlice - 1))
	indexIncrementForCrossPage = uint(rng.Intn(n))<<1 + intsPerPage + 1
}

////////////////////////////
// Random data generation //
////////////////////////////

var numbers []int

func init() {
	numbers = make([]int, intsPerSlice)
	for i, _ := range numbers {
		numbers[i] = rng.Int()
	}
}

/////////////
// Exports //
/////////////

func randomIntsInTheFirstPage() func() int {
	index := uint(0)
	return func() int {
		index = (index + indexIncrementForSinglePage) % intsPerPage
		return numbers[index]
	}
}

func randomIntsInPage(pageIndex uint) func() int {
	low := pageIndex * intsPerPage
	high := low + intsPerPage
	max := high
	numbersInPage := numbers[low:high:max]
	index := uint(0)
	return func() int {
		index = (index + indexIncrementForSinglePage) % intsPerPage
		return numbersInPage[index]
	}
}

func randomInts() func() int {
	index := uint(0)
	return func() int {
		index = (index + indexIncrementForCrossPage) % intsPerSlice
		return numbers[index]
	}
}

/////////////////
// Calibration //
/////////////////

func BenchmarkCalibrateEmptyFunction(b *testing.B) {
	for i, count := 0, b.N; i < count; i++ {
	}
}

func BenchmarkCalibrateRandomIntsInTheFirstPage(b *testing.B) {
	gen := randomIntsInTheFirstPage()
	for i, count := 0, b.N; i < count; i++ {
		gen()
	}
}

func BenchmarkCalibrateRandomIntsInTheFirstPageAlt(b *testing.B) {
	gen := randomIntsInPage(0)
	for i, count := 0, b.N; i < count; i++ {
		gen()
	}
}

func BenchmarkCalibrateRandomInts(b *testing.B) {
	gen := randomInts()
	for i, count := 0, b.N; i < count; i++ {
		gen()
	}
}

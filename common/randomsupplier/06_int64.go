package randomsupplier

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Int64 is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-6
type Int64 struct {
	index     uint
	increment uint
	count     uint
	numbers   []int64
}

// NewInt64 allocates and returns a new Int64.
func NewInt64() *Int64 {
	count := common.PageSizeInBytes >> 3
	gen := &Int64{
		count:   count,
		numbers: make([]int64, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Int64) Next() int64 {
	gen.index += gen.increment
	if gen.index >= gen.count {
		gen.index -= gen.count
	}
	return gen.numbers[gen.index]
}

// Reset reverts the state of gen to the same state as
// before Next() is called for the first time.
// It should be called every time when a new benchmark starts,
// before Next() is called for the first time.
func (gen *Int64) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Int64) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Int64SliceAsByteSlice(gen.numbers))
}

package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Int64Generator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-6
type Int64Generator struct {
	index     uint
	increment uint
	count     uint
	numbers   []int64
}

// NewInt64Generator allocates and returns a new Int64Generator.
func NewInt64Generator() *Int64Generator {
	count := common.PageSizeInBytes >> 3
	gen := &Int64Generator{
		count:   count,
		numbers: make([]int64, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Int64Generator) Next() int64 {
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
func (gen *Int64Generator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Int64Generator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Int64SliceAsByteSlice(gen.numbers))
}

package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Int32Generator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-5
type Int32Generator struct {
	index     uint
	increment uint
	count     uint
	numbers   []int32
}

// NewInt32Generator allocates and returns a new Int32Generator.
func NewInt32Generator() *Int32Generator {
	count := common.PageSizeInBytes >> 2
	gen := &Int32Generator{
		count:   count,
		numbers: make([]int32, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Int32Generator) Next() int32 {
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
func (gen *Int32Generator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Int32Generator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Int32SliceAsByteSlice(gen.numbers))
}

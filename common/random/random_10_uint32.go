package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint32Generator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-10
type Uint32Generator struct {
	index     uint
	increment uint
	count     uint
	numbers   []uint32
}

// NewUint32Generator allocates and returns a new Uint32Generator.
func NewUint32Generator() *Uint32Generator {
	count := common.PageSizeInBytes >> 2
	gen := &Uint32Generator{
		count:   count,
		numbers: make([]uint32, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint32Generator) Next() uint32 {
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
func (gen *Uint32Generator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint32Generator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Uint32SliceAsByteSlice(gen.numbers))
}

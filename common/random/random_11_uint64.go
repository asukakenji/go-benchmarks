package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint64Generator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-11
type Uint64Generator struct {
	index     uint
	increment uint
	count     uint
	numbers   []uint64
}

// NewUint64Generator allocates and returns a new Uint64Generator.
func NewUint64Generator() *Uint64Generator {
	count := common.PageSizeInBytes >> 3
	gen := &Uint64Generator{
		count:   count,
		numbers: make([]uint64, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint64Generator) Next() uint64 {
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
func (gen *Uint64Generator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint64Generator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Uint64SliceAsByteSlice(gen.numbers))
}

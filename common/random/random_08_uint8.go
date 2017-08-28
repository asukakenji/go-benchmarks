package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint8Generator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-8
type Uint8Generator struct {
	index     uint
	increment uint
	numbers   []uint8
}

// NewUint8Generator allocates and returns a new Uint8Generator.
func NewUint8Generator() *Uint8Generator {
	gen := &Uint8Generator{
		numbers: make([]uint8, common.PageSizeInBytes),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint8Generator) Next() uint8 {
	gen.index += gen.increment
	if gen.index >= common.PageSizeInBytes {
		gen.index -= common.PageSizeInBytes
	}
	return gen.numbers[gen.index]
}

// Reset reverts the state of gen to the same state as
// before Next() is called for the first time.
// It should be called every time when a new benchmark starts,
// before Next() is called for the first time.
func (gen *Uint8Generator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint8Generator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(common.PageSizeInBytes >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Uint8SliceAsByteSlice(gen.numbers))
}

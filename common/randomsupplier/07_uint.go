package randomsupplier

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-7
type Uint struct {
	index     uint
	increment uint
	numbers   []uint
}

// NewUint allocates and returns a new Uint.
func NewUint() *Uint {
	gen := &Uint{
		numbers: make([]uint, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint) Next() uint {
	gen.index += gen.increment
	if gen.index >= common.IntCountPerPage {
		gen.index -= common.IntCountPerPage
	}
	return gen.numbers[gen.index]
}

// Reset reverts the state of gen to the same state as
// before Next() is called for the first time.
// It should be called every time when a new benchmark starts,
// before Next() is called for the first time.
func (gen *Uint) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	// If IntCountPerPage is 512, we would like to generate an odd integer in
	// the range [1, 511], so that it always "generates" (in mathematical sense)
	// all indexes in modulo 512.
	// Since any odd integer in [1, 511] is any odd integer in [0, 510] + 1,
	// and that is any integer in [0, 255] * 2 + 1, we have:
	n := int(common.IntCountPerPage >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.UintSliceAsByteSlice(gen.numbers))
}

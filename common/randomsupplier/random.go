package randomsupplier

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// UintSlice is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-23(7)
type UintSlice struct {
	indexGen *Uint
	numbers  []uint
}

// NewUintSlice allocates and returns a new UintSlice.
func NewUintSlice() *UintSlice {
	gen := &UintSlice{
		indexGen: NewUint(),
		numbers:  make([]uint, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number slice.
func (gen *UintSlice) Next() []uint {
	a := gen.indexGen.Next() % common.IntCountPerPage
	b := gen.indexGen.Next() % common.IntCountPerPage
	if a > b {
		return gen.numbers[b:a]
	}
	return gen.numbers[a:b]
}

// Reset reverts the state of gen to the same state as
// before Next() is called for the first time.
// It should be called every time when a new benchmark starts,
// before Next() is called for the first time.
func (gen *UintSlice) Reset() {
	gen.indexGen.Reset()
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *UintSlice) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	rng.Read(reinterpret.UintSliceAsByteSlice(gen.numbers))

	gen.indexGen.Reinitialize()
}

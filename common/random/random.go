package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// UintSliceGenerator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-23(7)
type UintSliceGenerator struct {
	indexGen *UintGenerator
	numbers  []uint
}

// NewUintSliceGenerator allocates and returns a new UintSliceGenerator.
func NewUintSliceGenerator() *UintSliceGenerator {
	gen := &UintSliceGenerator{
		indexGen: NewUintGenerator(),
		numbers:  make([]uint, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number slice.
func (gen *UintSliceGenerator) Next() []uint {
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
func (gen *UintSliceGenerator) Reset() {
	gen.indexGen.Reset()
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *UintSliceGenerator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	rng.Read(reinterpret.UintSliceAsByteSlice(gen.numbers))

	gen.indexGen.Reinitialize()
}

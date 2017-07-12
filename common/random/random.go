package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// IntGenerator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-2
type IntGenerator struct {
	index     uint
	increment uint
	numbers   []int
}

// NewIntGenerator allocates and returns a new IntGenerator.
func NewIntGenerator() *IntGenerator {
	gen := &IntGenerator{
		numbers: make([]int, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *IntGenerator) Next() int {
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
func (gen *IntGenerator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *IntGenerator) Reinitialize() {
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
	rng.Read(reinterpret.IntSliceAsByteSlice(gen.numbers))
}

// UintGenerator is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-7
type UintGenerator struct {
	index     uint
	increment uint
	numbers   []uint
}

// NewUintGenerator allocates and returns a new UintGenerator.
func NewUintGenerator() *UintGenerator {
	gen := &UintGenerator{
		numbers: make([]uint, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *UintGenerator) Next() uint {
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
func (gen *UintGenerator) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *UintGenerator) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(common.IntCountPerPage >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.UintSliceAsByteSlice(gen.numbers))
}

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

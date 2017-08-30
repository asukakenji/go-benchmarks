package randomsupplier

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint64 is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-11
type Uint64 struct {
	index     uint
	increment uint
	count     uint
	numbers   []uint64
}

// NewUint64 allocates and returns a new Uint64.
func NewUint64() *Uint64 {
	count := common.PageSizeInBytes >> 3
	gen := &Uint64{
		count:   count,
		numbers: make([]uint64, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint64) Next() uint64 {
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
func (gen *Uint64) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint64) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Uint64SliceAsByteSlice(gen.numbers))
}

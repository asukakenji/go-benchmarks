package randomsupplier

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

// Uint32 is a type for generating reproducible random numbers
// used in test cases or benchmarks.
//
// ID: RNG-10
type Uint32 struct {
	index     uint
	increment uint
	count     uint
	numbers   []uint32
}

// NewUint32 allocates and returns a new Uint32.
func NewUint32() *Uint32 {
	count := common.PageSizeInBytes >> 2
	gen := &Uint32{
		count:   count,
		numbers: make([]uint32, count),
	}
	gen.Reinitialize()
	return gen
}

// Next returns the next random number.
func (gen *Uint32) Next() uint32 {
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
func (gen *Uint32) Reset() {
	gen.index = 0
}

// Reinitialize generates a new set of random numbers in gen.
func (gen *Uint32) Reinitialize() {
	seed := time.Now().UTC().UnixNano()
	src := rand.NewSource(seed)
	rng := rand.New(src)

	n := int(gen.count >> 1)
	gen.index = 0
	gen.increment = uint(rng.Intn(n))<<1 + 1
	rng.Read(reinterpret.Uint32SliceAsByteSlice(gen.numbers))
}

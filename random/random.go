package random

import (
	"math/rand"
	"time"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/reinterpret"
)

type IntGenerator struct {
	index     uint
	increment uint
	numbers   []int
}

func NewIntGenerator() *IntGenerator {
	gen := &IntGenerator{
		numbers: make([]int, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

func (gen *IntGenerator) Next() int {
	gen.index += gen.increment
	if gen.index >= common.IntCountPerPage {
		gen.index -= common.IntCountPerPage
	}
	return gen.numbers[gen.index]
}

func (gen *IntGenerator) Reset() {
	gen.index = 0
}

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

type UintGenerator struct {
	index     uint
	increment uint
	numbers   []uint
}

func NewUintGenerator() *UintGenerator {
	gen := &UintGenerator{
		numbers: make([]uint, common.IntCountPerPage),
	}
	gen.Reinitialize()
	return gen
}

func (gen *UintGenerator) Next() uint {
	gen.index += gen.increment
	if gen.index >= common.IntCountPerPage {
		gen.index -= common.IntCountPerPage
	}
	return gen.numbers[gen.index]
}

func (gen *UintGenerator) Reset() {
	gen.index = 0
}

func (gen *UintGenerator) Reinitialize() {
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

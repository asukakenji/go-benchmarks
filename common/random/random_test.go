package random_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/random"
)

func TestNewIntGeneratorNext(t *testing.T) {
	gen := random.NewIntGenerator()
Outer:
	for {
		for i := uint(0); i < common.IntCountPerPage; i++ {
			n := gen.Next()
			if n < 0 {
				break Outer
			}
		}
		gen.Reinitialize()
	}
	t.Log("IntGenerator can generate an int < 0")

	gen.Reset()
	saved := []int{}
	for i := uint(0); i < common.IntCountPerPage; i++ {
		saved = append(saved, gen.Next())
	}
	for i, expected := range saved {
		got := gen.Next()
		if got != expected {
			t.Errorf("gen.Next() = %d at iteration %d, expected %d", got, i, expected)
		}
	}
}

func TestNewUintGeneratorNext(t *testing.T) {
	gen := random.NewUintGenerator()
Outer:
	for {
		for i := uint(0); i < common.IntCountPerPage; i++ {
			n := gen.Next()
			if n >= 0x8000000000000000 {
				break Outer
			}
		}
		gen.Reinitialize()
	}
	t.Log("UintGenerator can generate an uint >= 0x8000000000000000")

	gen.Reset()
	saved := []uint{}
	for i := uint(0); i < common.IntCountPerPage; i++ {
		saved = append(saved, gen.Next())
	}
	for i, expected := range saved {
		got := gen.Next()
		if got != expected {
			t.Errorf("gen.Next() = %d at iteration %d, expected %d", got, i, expected)
		}
	}
}

func TestNewUint32GeneratorNext(t *testing.T) {
	gen := random.NewUint32Generator()
Outer:
	for {
		for i, count := uint(0), common.PageSizeInBytes>>2; i < count; i++ {
			n := gen.Next()
			if n >= 0x80000000 {
				break Outer
			}
		}
		gen.Reinitialize()
	}
	t.Log("Uint32Generator can generate an uint32 >= 0x80000000")

	gen.Reset()
	saved := []uint32{}
	for i, count := uint(0), common.PageSizeInBytes>>2; i < count; i++ {
		saved = append(saved, gen.Next())
	}
	for i, expected := range saved {
		got := gen.Next()
		if got != expected {
			t.Errorf("gen.Next() = %d at iteration %d, expected %d", got, i, expected)
		}
	}
}

func TestNewUint64GeneratorNext(t *testing.T) {
	gen := random.NewUint64Generator()
Outer:
	for {
		for i, count := uint(0), common.PageSizeInBytes>>3; i < count; i++ {
			n := gen.Next()
			if n >= 0x8000000000000000 {
				break Outer
			}
		}
		gen.Reinitialize()
	}
	t.Log("Uint64Generator can generate an uint64 >= 0x8000000000000000")

	gen.Reset()
	saved := []uint64{}
	for i, count := uint(0), common.PageSizeInBytes>>3; i < count; i++ {
		saved = append(saved, gen.Next())
	}
	for i, expected := range saved {
		got := gen.Next()
		if got != expected {
			t.Errorf("gen.Next() = %d at iteration %d, expected %d", got, i, expected)
		}
	}
}

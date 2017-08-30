package randomsupplier_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/common/randomsupplier"
)

func TestInt(t *testing.T) {
	gen := randomsupplier.NewInt()
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
	t.Log("Int can generate an int < 0")

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

func TestUint(t *testing.T) {
	gen := randomsupplier.NewUint()
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
	t.Log("Uint can generate an uint >= 0x8000000000000000")

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

func TestUint32(t *testing.T) {
	gen := randomsupplier.NewUint32()
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
	t.Log("Uint32 can generate an uint32 >= 0x80000000")

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

func TestUint64(t *testing.T) {
	gen := randomsupplier.NewUint64()
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
	t.Log("Uint64 can generate an uint64 >= 0x8000000000000000")

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

func TestUintSlice(t *testing.T) {
	gen := randomsupplier.NewUintSlice()
	for i, count := uint(0), uint(1024); i < count; i++ {
		for j, count := uint(0), uint(1024); j < count; j++ {
			gen.Next()
		}
		gen.Reinitialize()
	}
	gen.Reset()
}

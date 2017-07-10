package random_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/random"
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
	t.Log("NewIntGenerator can generator an int < 0")

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
	t.Log("NewUintGenerator can generator an uint >= 0x8000000000000000")
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
	t.Log("NewUintGenerator can generator an uint >= 0x8000000000000000")

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
	t.Log("NewUintGenerator can generator an uint >= 0x8000000000000000")
}

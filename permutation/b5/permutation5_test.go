package b5_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b5"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b5.Permutation5Inc", b5.Permutation5Inc},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

package b1a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b1a"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b1a.Permutation1ParamOrder0", b1a.Permutation1ParamOrder0},
		{"b1a.Permutation1ParamOrder1", b1a.Permutation1ParamOrder1},
		{"b1a.Permutation1ParamOrder2", b1a.Permutation1ParamOrder2},
		{"b1a.Permutation1ParamOrder3", b1a.Permutation1ParamOrder3},
		{"b1a.Permutation1ParamOrder4", b1a.Permutation1ParamOrder4},
		{"b1a.Permutation1ParamOrder5", b1a.Permutation1ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

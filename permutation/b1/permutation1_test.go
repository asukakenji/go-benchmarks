package b1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b1.Permutation1ParamOrder0", b1.Permutation1ParamOrder0},
		{"b1.Permutation1ParamOrder1", b1.Permutation1ParamOrder1},
		{"b1.Permutation1ParamOrder2", b1.Permutation1ParamOrder2},
		{"b1.Permutation1ParamOrder3", b1.Permutation1ParamOrder3},
		{"b1.Permutation1ParamOrder4", b1.Permutation1ParamOrder4},
		{"b1.Permutation1ParamOrder5", b1.Permutation1ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

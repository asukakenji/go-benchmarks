package b0_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b0.Permutation0ParamOrder0", b0.Permutation0ParamOrder0},
		{"b0.Permutation0ParamOrder1", b0.Permutation0ParamOrder1},
		{"b0.Permutation0ParamOrder2", b0.Permutation0ParamOrder2},
		{"b0.Permutation0ParamOrder3", b0.Permutation0ParamOrder3},
		{"b0.Permutation0ParamOrder4", b0.Permutation0ParamOrder4},
		{"b0.Permutation0ParamOrder5", b0.Permutation0ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

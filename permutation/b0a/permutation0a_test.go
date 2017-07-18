package b0a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0a"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b0a.Permutation0ParamOrder0", b0a.Permutation0ParamOrder0},
		{"b0a.Permutation0ParamOrder1", b0a.Permutation0ParamOrder1},
		{"b0a.Permutation0ParamOrder2", b0a.Permutation0ParamOrder2},
		{"b0a.Permutation0ParamOrder3", b0a.Permutation0ParamOrder3},
		{"b0a.Permutation0ParamOrder4", b0a.Permutation0ParamOrder4},
		{"b0a.Permutation0ParamOrder5", b0a.Permutation0ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

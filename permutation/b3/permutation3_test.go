package b3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b3"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b3.Permutation3ParamOrder0", b3.Permutation3ParamOrder0},
		{"b3.Permutation3ParamOrder1", b3.Permutation3ParamOrder1},
		{"b3.Permutation3ParamOrder2", b3.Permutation3ParamOrder2},
		{"b3.Permutation3ParamOrder3", b3.Permutation3ParamOrder3},
		{"b3.Permutation3ParamOrder4", b3.Permutation3ParamOrder4},
		{"b3.Permutation3ParamOrder5", b3.Permutation3ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

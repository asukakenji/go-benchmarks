package b3_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b3"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b3.Permutation3ParamOrder0", F: b3.Permutation3ParamOrder0},
		{Name: "b3.Permutation3ParamOrder1", F: b3.Permutation3ParamOrder1},
		{Name: "b3.Permutation3ParamOrder2", F: b3.Permutation3ParamOrder2},
		{Name: "b3.Permutation3ParamOrder3", F: b3.Permutation3ParamOrder3},
		{Name: "b3.Permutation3ParamOrder4", F: b3.Permutation3ParamOrder4},
		{Name: "b3.Permutation3ParamOrder5", F: b3.Permutation3ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

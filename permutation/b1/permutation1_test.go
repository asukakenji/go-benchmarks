package b1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b1.Permutation1ParamOrder0", F: b1.Permutation1ParamOrder0},
		{Name: "b1.Permutation1ParamOrder1", F: b1.Permutation1ParamOrder1},
		{Name: "b1.Permutation1ParamOrder2", F: b1.Permutation1ParamOrder2},
		{Name: "b1.Permutation1ParamOrder3", F: b1.Permutation1ParamOrder3},
		{Name: "b1.Permutation1ParamOrder4", F: b1.Permutation1ParamOrder4},
		{Name: "b1.Permutation1ParamOrder5", F: b1.Permutation1ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

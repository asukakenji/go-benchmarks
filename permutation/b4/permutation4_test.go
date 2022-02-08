package b4_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b4"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b4.Permutation4ParamOrder0", F: b4.Permutation4ParamOrder0},
		{Name: "b4.Permutation4ParamOrder1", F: b4.Permutation4ParamOrder1},
		{Name: "b4.Permutation4ParamOrder2", F: b4.Permutation4ParamOrder2},
		{Name: "b4.Permutation4ParamOrder3", F: b4.Permutation4ParamOrder3},
		{Name: "b4.Permutation4ParamOrder4", F: b4.Permutation4ParamOrder4},
		{Name: "b4.Permutation4ParamOrder5", F: b4.Permutation4ParamOrder5},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

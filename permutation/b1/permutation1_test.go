package b1_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b1"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b1.Permutation1ParamOrder0", F: b1.Permutation1ParamOrder0[uint]},
		{Name: "b1.Permutation1ParamOrder1", F: b1.Permutation1ParamOrder1[uint]},
		{Name: "b1.Permutation1ParamOrder2", F: b1.Permutation1ParamOrder2[uint]},
		{Name: "b1.Permutation1ParamOrder3", F: b1.Permutation1ParamOrder3[uint]},
		{Name: "b1.Permutation1ParamOrder4", F: b1.Permutation1ParamOrder4[uint]},
		{Name: "b1.Permutation1ParamOrder5", F: b1.Permutation1ParamOrder5[uint]},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

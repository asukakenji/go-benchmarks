package b0a_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0a"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b0a.Permutation0ParamOrder0", F: b0a.Permutation0ParamOrder0[uint]},
		{Name: "b0a.Permutation0ParamOrder1", F: b0a.Permutation0ParamOrder1[uint]},
		{Name: "b0a.Permutation0ParamOrder2", F: b0a.Permutation0ParamOrder2[uint]},
		{Name: "b0a.Permutation0ParamOrder3", F: b0a.Permutation0ParamOrder3[uint]},
		{Name: "b0a.Permutation0ParamOrder4", F: b0a.Permutation0ParamOrder4[uint]},
		{Name: "b0a.Permutation0ParamOrder5", F: b0a.Permutation0ParamOrder5[uint]},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

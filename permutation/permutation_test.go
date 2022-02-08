package permutation_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func adapter(x func(func([]uint), []uint)) func([]uint, func([]uint)) {
	return func(s []uint, f func([]uint)) {
		x(f, s)
	}
}

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "permutation.PermutationOrder0", F: permutation.PermutationOrder0[uint]},
		{Name: "permutation.PermutationOrder1", F: adapter(permutation.PermutationOrder1[uint])},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

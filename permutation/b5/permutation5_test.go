package b5_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b5"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b5.Permutation5Inc", F: b5.Permutation5Inc[uint]},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

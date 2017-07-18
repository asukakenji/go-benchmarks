package b2_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b2"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{"b2.Permutation2ParamOrder00", b2.Permutation2ParamOrder00},
		{"b2.Permutation2ParamOrder01", b2.Permutation2ParamOrder01},
		{"b2.Permutation2ParamOrder02", b2.Permutation2ParamOrder02},
		{"b2.Permutation2ParamOrder03", b2.Permutation2ParamOrder03},
		{"b2.Permutation2ParamOrder10", b2.Permutation2ParamOrder10},
		{"b2.Permutation2ParamOrder11", b2.Permutation2ParamOrder11},
		{"b2.Permutation2ParamOrder12", b2.Permutation2ParamOrder12},
		{"b2.Permutation2ParamOrder13", b2.Permutation2ParamOrder13},
		{"b2.Permutation2ParamOrder20", b2.Permutation2ParamOrder20},
		{"b2.Permutation2ParamOrder21", b2.Permutation2ParamOrder21},
		{"b2.Permutation2ParamOrder22", b2.Permutation2ParamOrder22},
		{"b2.Permutation2ParamOrder23", b2.Permutation2ParamOrder23},
		{"b2.Permutation2ParamOrder30", b2.Permutation2ParamOrder30},
		{"b2.Permutation2ParamOrder31", b2.Permutation2ParamOrder31},
		{"b2.Permutation2ParamOrder32", b2.Permutation2ParamOrder32},
		{"b2.Permutation2ParamOrder33", b2.Permutation2ParamOrder33},
		{"b2.Permutation2ParamOrder40", b2.Permutation2ParamOrder40},
		{"b2.Permutation2ParamOrder41", b2.Permutation2ParamOrder41},
		{"b2.Permutation2ParamOrder42", b2.Permutation2ParamOrder42},
		{"b2.Permutation2ParamOrder43", b2.Permutation2ParamOrder43},
		{"b2.Permutation2ParamOrder50", b2.Permutation2ParamOrder50},
		{"b2.Permutation2ParamOrder51", b2.Permutation2ParamOrder51},
		{"b2.Permutation2ParamOrder52", b2.Permutation2ParamOrder52},
		{"b2.Permutation2ParamOrder53", b2.Permutation2ParamOrder53},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

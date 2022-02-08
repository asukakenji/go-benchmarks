package b2_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b2"
	"github.com/asukakenji/go-benchmarks/permutation/tcommon"
)

func TestPermutation(t *testing.T) {
	implementations := []tcommon.Implementation{
		{Name: "b2.Permutation2ParamOrder00", F: b2.Permutation2ParamOrder00},
		{Name: "b2.Permutation2ParamOrder01", F: b2.Permutation2ParamOrder01},
		{Name: "b2.Permutation2ParamOrder02", F: b2.Permutation2ParamOrder02},
		{Name: "b2.Permutation2ParamOrder03", F: b2.Permutation2ParamOrder03},
		{Name: "b2.Permutation2ParamOrder10", F: b2.Permutation2ParamOrder10},
		{Name: "b2.Permutation2ParamOrder11", F: b2.Permutation2ParamOrder11},
		{Name: "b2.Permutation2ParamOrder12", F: b2.Permutation2ParamOrder12},
		{Name: "b2.Permutation2ParamOrder13", F: b2.Permutation2ParamOrder13},
		{Name: "b2.Permutation2ParamOrder20", F: b2.Permutation2ParamOrder20},
		{Name: "b2.Permutation2ParamOrder21", F: b2.Permutation2ParamOrder21},
		{Name: "b2.Permutation2ParamOrder22", F: b2.Permutation2ParamOrder22},
		{Name: "b2.Permutation2ParamOrder23", F: b2.Permutation2ParamOrder23},
		{Name: "b2.Permutation2ParamOrder30", F: b2.Permutation2ParamOrder30},
		{Name: "b2.Permutation2ParamOrder31", F: b2.Permutation2ParamOrder31},
		{Name: "b2.Permutation2ParamOrder32", F: b2.Permutation2ParamOrder32},
		{Name: "b2.Permutation2ParamOrder33", F: b2.Permutation2ParamOrder33},
		{Name: "b2.Permutation2ParamOrder40", F: b2.Permutation2ParamOrder40},
		{Name: "b2.Permutation2ParamOrder41", F: b2.Permutation2ParamOrder41},
		{Name: "b2.Permutation2ParamOrder42", F: b2.Permutation2ParamOrder42},
		{Name: "b2.Permutation2ParamOrder43", F: b2.Permutation2ParamOrder43},
		{Name: "b2.Permutation2ParamOrder50", F: b2.Permutation2ParamOrder50},
		{Name: "b2.Permutation2ParamOrder51", F: b2.Permutation2ParamOrder51},
		{Name: "b2.Permutation2ParamOrder52", F: b2.Permutation2ParamOrder52},
		{Name: "b2.Permutation2ParamOrder53", F: b2.Permutation2ParamOrder53},
	}
	for _, impl := range implementations {
		tcommon.TestPermutation(t, impl)
	}
}

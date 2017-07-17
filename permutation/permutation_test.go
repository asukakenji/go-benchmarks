package permutation_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation/b0"
	"github.com/asukakenji/go-benchmarks/permutation/b0a"
	"github.com/asukakenji/go-benchmarks/permutation/b1"
)

func TestPermutation(t *testing.T) {
	implementations := []struct {
		name string
		f    func([]uint, func([]uint))
	}{
		{"b0.Permutation0ParamOrder0", b0.Permutation0ParamOrder0},
		{"b0.Permutation0ParamOrder1", b0.Permutation0ParamOrder1},
		{"b0.Permutation0ParamOrder2", b0.Permutation0ParamOrder2},
		{"b0.Permutation0ParamOrder3", b0.Permutation0ParamOrder3},
		{"b0.Permutation0ParamOrder4", b0.Permutation0ParamOrder4},
		{"b0.Permutation0ParamOrder5", b0.Permutation0ParamOrder5},
		{"b0a.Permutation0ParamOrder0", b0a.Permutation0ParamOrder0},
		{"b0a.Permutation0ParamOrder1", b0a.Permutation0ParamOrder1},
		{"b0a.Permutation0ParamOrder2", b0a.Permutation0ParamOrder2},
		{"b0a.Permutation0ParamOrder3", b0a.Permutation0ParamOrder3},
		{"b0a.Permutation0ParamOrder4", b0a.Permutation0ParamOrder4},
		{"b0a.Permutation0ParamOrder5", b0a.Permutation0ParamOrder5},
		{"b1.Permutation1ParamOrder0", b1.Permutation1ParamOrder0},
		{"b1.Permutation1ParamOrder1", b1.Permutation1ParamOrder1},
		{"b1.Permutation1ParamOrder2", b1.Permutation1ParamOrder2},
		{"b1.Permutation1ParamOrder3", b1.Permutation1ParamOrder3},
		{"b1.Permutation1ParamOrder4", b1.Permutation1ParamOrder4},
		{"b1.Permutation1ParamOrder5", b1.Permutation1ParamOrder5},
	}
	cases := []struct {
		s        []uint
		expected [][]uint
	}{
		{
			[]uint{},
			[][]uint{
				{},
			},
		},
		{
			[]uint{1},
			[][]uint{
				{1},
			},
		},
		{
			[]uint{1, 2},
			[][]uint{
				{1, 2},
				{2, 1},
			},
		},
		{
			[]uint{1, 2, 3},
			[][]uint{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
		{
			[]uint{1, 2, 3, 4},
			[][]uint{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 3, 2},
				{1, 4, 2, 3},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 3, 1},
				{2, 4, 1, 3},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 2, 3, 1},
				{4, 2, 1, 3},
				{4, 3, 2, 1},
				{4, 3, 1, 2},
				{4, 1, 3, 2},
				{4, 1, 2, 3},
			},
		},
	}
	for _, impl := range implementations {
		for _, c := range cases {
			sCopy := make([]uint, len(c.s))
			copy(sCopy, c.s)
			got := [][]uint{}
			f := func(x []uint) {
				xCopy := make([]uint, len(x))
				copy(xCopy, x)
				got = append(got, xCopy)
			}
			impl.f(c.s, f)
			if !reflect.DeepEqual(c.s, sCopy) {
				t.Errorf("%s changed the input: %v, expected %v", impl.name, c.s, sCopy)
			}
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s(%v, <func>) = %v, expected %v", impl.name, c.s, got, c.expected)
			}
		}
	}
}

package permutation_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation"
)

func TestPermutation(t *testing.T) {
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
	for _, c := range cases {
		got := [][]uint{}
		f := func(x []uint) {
			x2 := make([]uint, len(x))
			copy(x2, x)
			got = append(got, x2)
		}
		permutation.Permutation0(c.s, f)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("Permutation0(%d) = %v, expected %v", c.s, got, c.expected)
		}
	}
}

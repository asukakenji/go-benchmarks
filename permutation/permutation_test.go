package permutation_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/permutation"
)

func TestPermutation(t *testing.T) {
	implementations := []struct {
		name string
		f    func([]uint, func([]uint))
	}{
		{"Permutation0", permutation.Permutation0},
		{"Permutation0Alt1", permutation.Permutation0Alt1},
		{"Permutation1Alt1", permutation.Permutation1Alt1},
		{"Permutation1Alt2", permutation.Permutation1Alt2},
		{"Permutation1Alt3", permutation.Permutation1Alt3},
		{"Permutation1Alt4", permutation.Permutation1Alt4},
		{"Permutation1Alt5", permutation.Permutation1Alt5},
		{"Permutation1Alt6", permutation.Permutation1Alt6},
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

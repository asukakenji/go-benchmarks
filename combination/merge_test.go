package combination_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
	"github.com/asukakenji/go-benchmarks/common/benchmark"
)

func TestMergeEachPairOfElementsInMultisetByAddition(t *testing.T) {
	implementations := []struct {
		name string
		f    func([]uint, func([]uint))
	}{
		{"MergeEachPairOfElementsInMultisetByAddition0", combination.MergeEachPairOfElementsInMultisetByAddition0},
		{"MergeEachPairOfElementsInMultisetByAddition1", combination.MergeEachPairOfElementsInMultisetByAddition1},
		{"MergeEachPairOfElementsInMultisetByAddition2Alt1", combination.MergeEachPairOfElementsInMultisetByAddition2Alt1},
		{"MergeEachPairOfElementsInMultisetByAddition2Alt2", combination.MergeEachPairOfElementsInMultisetByAddition2Alt2},
		{"MergeEachPairOfElementsInMultisetByAddition3", combination.MergeEachPairOfElementsInMultisetByAddition3},
	}
	cases := []struct {
		set      []uint
		expected [][]uint
	}{
		{
			[]uint{},
			[][]uint{},
		},
		{
			[]uint{1},
			[][]uint{},
		},
		{
			[]uint{1, 2},
			[][]uint{
				{2 + 1},
			},
		},
		{
			[]uint{1, 2, 3},
			[][]uint{
				{2 + 1, 3},
				{2, 3 + 1},
				{1, 3 + 2},
			},
		},
		{
			[]uint{1, 2, 3, 4},
			[][]uint{
				{2 + 1, 3, 4},
				{2, 3 + 1, 4},
				{2, 3, 4 + 1},
				{1, 3 + 2, 4},
				{1, 3, 4 + 2},
				{2, 1, 4 + 3},
			},
		},
		{
			[]uint{1, 2, 3, 4, 5},
			[][]uint{
				{2 + 1, 3, 4, 5},
				{2, 3 + 1, 4, 5},
				{2, 3, 4 + 1, 5},
				{2, 3, 4, 5 + 1},
				{1, 3 + 2, 4, 5},
				{1, 3, 4 + 2, 5},
				{1, 3, 4, 5 + 2},
				{2, 1, 4 + 3, 5},
				{2, 1, 4, 5 + 3},
				{2, 3, 1, 5 + 4},
			},
		},
	}
	for _, impl := range implementations {
		for _, c := range cases {
			setCopy := make([]uint, len(c.set))
			copy(setCopy, c.set)
			got := [][]uint{}
			f := func(x []uint) {
				xCopy := make([]uint, len(x))
				copy(xCopy, x)
				got = append(got, xCopy)
			}
			impl.f(c.set, f)
			if !reflect.DeepEqual(c.set, setCopy) {
				t.Errorf("%s changed the input: %v, expected %v", impl.name, c.set, setCopy)
			}
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s(%v, <func>) = %v, expected %v", impl.name, c.set, got, c.expected)
			}
		}
	}
}

// BenchmarkMergeEachPairOfElementsInMultisetByAddition0-8       	   20000	     81051 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition1-8       	   20000	     83577 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt1-8   	   20000	     80658 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt2-8   	   20000	     79481 ns/op <- Best
// BenchmarkMergeEachPairOfElementsInMultisetByAddition3-8       	   20000	     87955 ns/op

func BenchmarkMergeEachPairOfElementsInMultisetByAddition0(b *testing.B) {
	benchmark.UintSliceGeneratorWithRandom(b, combination.MergeEachPairOfElementsInMultisetByAddition0)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition1(b *testing.B) {
	benchmark.UintSliceGeneratorWithRandom(b, combination.MergeEachPairOfElementsInMultisetByAddition1)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt1(b *testing.B) {
	benchmark.UintSliceGeneratorWithRandom(b, combination.MergeEachPairOfElementsInMultisetByAddition2Alt1)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt2(b *testing.B) {
	benchmark.UintSliceGeneratorWithRandom(b, combination.MergeEachPairOfElementsInMultisetByAddition2Alt2)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition3(b *testing.B) {
	benchmark.UintSliceGeneratorWithRandom(b, combination.MergeEachPairOfElementsInMultisetByAddition3)
}

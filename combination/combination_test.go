package combination_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
)

type Collector [][]uint

func NewCollector() *Collector {
	collector := Collector([][]uint{})
	return &collector
}

func (c *Collector) Collect(x []uint) {
	xCopy := make([]uint, len(x))
	copy(xCopy, x)
	*c = append(*c, xCopy)
}

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
			collector := NewCollector()
			impl.f(c.set, collector.Collect)
			got := ([][]uint)(*collector)
			if !reflect.DeepEqual(got, c.expected) {
				t.Errorf("%s(%v, collector.Collect) = %v, expected %v", impl.name, c.set, got, c.expected)
			}
		}
	}
}

var benchmarkCases = [][]uint{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	{0, 0, 1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0},
	{0, 0, 1, 1, 0, 1, 2, 0, 1, 3, 0, 1, 4, 0, 1, 5, 0, 1, 6, 0, 1, 7, 0, 1, 8, 0, 1, 9, 0, 1},
	{0, 0, 1, 9, 1, 0, 1, 9, 2, 0, 1, 9, 3, 0, 1, 9, 4, 0, 1, 9, 5, 0, 1, 9, 6, 0, 1, 9, 7, 0, 1, 9, 8, 0, 1, 9, 9, 0, 1, 9},
}

func dummy(_ []uint) {
}

// BenchmarkMergeEachPairOfElementsInMultisetByAddition0-8       	 1000000	      1195 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition1-8       	 1000000	      1226 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt1-8   	 1000000	      1157 ns/op
// BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt2-8   	 1000000	      1130 ns/op <- Best
// BenchmarkMergeEachPairOfElementsInMultisetByAddition3-8       	 1000000	      1254 ns/op

func benchmarkMergeEachPairOfElementsInMultisetByAddition(b *testing.B, impl func([]uint, func([]uint))) {
	n := len(benchmarkCases)
	index := 0
	for i, count := 0, b.N; i < count; i++ {
		impl(benchmarkCases[index], dummy)
		index++
		if index == n {
			index = 0
		}
	}
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition0(b *testing.B) {
	benchmarkMergeEachPairOfElementsInMultisetByAddition(b, combination.MergeEachPairOfElementsInMultisetByAddition0)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition1(b *testing.B) {
	benchmarkMergeEachPairOfElementsInMultisetByAddition(b, combination.MergeEachPairOfElementsInMultisetByAddition1)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt1(b *testing.B) {
	benchmarkMergeEachPairOfElementsInMultisetByAddition(b, combination.MergeEachPairOfElementsInMultisetByAddition2Alt1)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition2Alt2(b *testing.B) {
	benchmarkMergeEachPairOfElementsInMultisetByAddition(b, combination.MergeEachPairOfElementsInMultisetByAddition2Alt2)
}

func BenchmarkMergeEachPairOfElementsInMultisetByAddition3(b *testing.B) {
	benchmarkMergeEachPairOfElementsInMultisetByAddition(b, combination.MergeEachPairOfElementsInMultisetByAddition3)
}

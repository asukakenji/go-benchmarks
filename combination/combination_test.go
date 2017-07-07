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

func TestMergeEachPairOfElementsInSetByAddition(t *testing.T) {
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
	for _, c := range cases {
		collector := NewCollector()
		combination.MergeEachPairOfElementsInSetByAddition(c.set, collector.Collect)
		got := ([][]uint)(*collector)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("MergeEachPairOfElementsInSetByAddition(%v, collector.Collect) = %v, expected %v", c.set, got, c.expected)
		}
	}
}

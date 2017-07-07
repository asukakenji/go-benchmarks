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
				{3},
			},
		},
		{
			[]uint{1, 2, 3},
			[][]uint{
				{3, 3},
				{2, 4},
				{1, 5},
			},
		},
		{
			[]uint{1, 2, 3, 4},
			[][]uint{
				{3, 3, 4},
				{2, 4, 4},
				{2, 3, 5},
				{1, 5, 4},
				{1, 3, 6},
				{2, 1, 7},
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

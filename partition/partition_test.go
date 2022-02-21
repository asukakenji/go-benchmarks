package partition_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/partition"
)

func testPartition(t *testing.T, partition func(int) [][]int) {
	cases := []struct {
		n        int
		expected [][]int
	}{
		{1, [][]int{
			{1},
		}},
		{2, [][]int{
			{1, 1},
			{2},
		}},
		{3, [][]int{
			{1, 1, 1},
			{2, 1},
			{3},
		}},
		{4, [][]int{
			{1, 1, 1, 1},
			{2, 1, 1}, {2, 2},
			{3, 1},
			{4},
		}},
		{5, [][]int{
			{1, 1, 1, 1, 1},
			{2, 1, 1, 1}, {2, 2, 1},
			{3, 1, 1}, {3, 2},
			{4, 1},
			{5},
		}},
		{6, [][]int{
			{1, 1, 1, 1, 1, 1},
			{2, 1, 1, 1, 1}, {2, 2, 1, 1}, {2, 2, 2},
			{3, 1, 1, 1}, {3, 2, 1}, {3, 3},
			{4, 1, 1}, {4, 2},
			{5, 1},
			{6},
		}},
		{7, [][]int{
			{1, 1, 1, 1, 1, 1, 1},
			{2, 1, 1, 1, 1, 1}, {2, 2, 1, 1, 1}, {2, 2, 2, 1},
			{3, 1, 1, 1, 1}, {3, 2, 1, 1}, {3, 2, 2}, {3, 3, 1},
			{4, 1, 1, 1}, {4, 2, 1}, {4, 3},
			{5, 1, 1}, {5, 2},
			{6, 1},
			{7},
		}},
		{8, [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{2, 1, 1, 1, 1, 1, 1}, {2, 2, 1, 1, 1, 1}, {2, 2, 2, 1, 1}, {2, 2, 2, 2},
			{3, 1, 1, 1, 1, 1}, {3, 2, 1, 1, 1}, {3, 2, 2, 1}, {3, 3, 1, 1}, {3, 3, 2},
			{4, 1, 1, 1, 1}, {4, 2, 1, 1}, {4, 2, 2}, {4, 3, 1}, {4, 4},
			{5, 1, 1, 1}, {5, 2, 1}, {5, 3},
			{6, 1, 1}, {6, 2},
			{7, 1},
			{8},
		}},
		{9, [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{2, 1, 1, 1, 1, 1, 1, 1}, {2, 2, 1, 1, 1, 1, 1}, {2, 2, 2, 1, 1, 1}, {2, 2, 2, 2, 1},
			{3, 1, 1, 1, 1, 1, 1}, {3, 2, 1, 1, 1, 1}, {3, 2, 2, 1, 1}, {3, 2, 2, 2}, {3, 3, 1, 1, 1}, {3, 3, 2, 1}, {3, 3, 3},
			{4, 1, 1, 1, 1, 1}, {4, 2, 1, 1, 1}, {4, 2, 2, 1}, {4, 3, 1, 1}, {4, 3, 2}, {4, 4, 1},
			{5, 1, 1, 1, 1}, {5, 2, 1, 1}, {5, 2, 2}, {5, 3, 1}, {5, 4},
			{6, 1, 1, 1}, {6, 2, 1}, {6, 3},
			{7, 1, 1}, {7, 2},
			{8, 1},
			{9},
		}},
		{10, [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{2, 1, 1, 1, 1, 1, 1, 1, 1}, {2, 2, 1, 1, 1, 1, 1, 1}, {2, 2, 2, 1, 1, 1, 1}, {2, 2, 2, 2, 1, 1}, {2, 2, 2, 2, 2},
			{3, 1, 1, 1, 1, 1, 1, 1}, {3, 2, 1, 1, 1, 1, 1}, {3, 2, 2, 1, 1, 1}, {3, 2, 2, 2, 1}, {3, 3, 1, 1, 1, 1}, {3, 3, 2, 1, 1}, {3, 3, 2, 2}, {3, 3, 3, 1},
			{4, 1, 1, 1, 1, 1, 1}, {4, 2, 1, 1, 1, 1}, {4, 2, 2, 1, 1}, {4, 2, 2, 2}, {4, 3, 1, 1, 1}, {4, 3, 2, 1}, {4, 3, 3}, {4, 4, 1, 1}, {4, 4, 2},
			{5, 1, 1, 1, 1, 1}, {5, 2, 1, 1, 1}, {5, 2, 2, 1}, {5, 3, 1, 1}, {5, 3, 2}, {5, 4, 1}, {5, 5},
			{6, 1, 1, 1, 1}, {6, 2, 1, 1}, {6, 2, 2}, {6, 3, 1}, {6, 4},
			{7, 1, 1, 1}, {7, 2, 1}, {7, 3},
			{8, 1, 1}, {8, 2},
			{9, 1},
			{10},
		}},
	}
	for _, c := range cases {
		got := partition(c.n)
		if !reflect.DeepEqual(got, c.expected) {
			t.Errorf("Partition(%d) = %v, expected: %v", c.n, got, c.expected)
		}
	}
}

func TestPartition1(t *testing.T) {
	testPartition(t, partition.Partition1)
}

func TestPartition2(t *testing.T) {
	testPartition(t, partition.Partition2)
}

func TestPartition2a(t *testing.T) {
	testPartition(t, partition.Partition2a)
}

func TestPartition3(t *testing.T) {
	testPartition(t, partition.Partition3)
}
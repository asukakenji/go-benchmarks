package combination_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
)

// TODO: Make X = -1 and fix other implementations
const X = 0

func TestCombinationCount(t *testing.T) {
	cases := []struct {
		n        int
		r        int
		expected int
	}{
		{-2, -2, X}, {-2, -1, X}, {-2, 0, X}, {-2, 1, X}, {-2, 2, X}, {-2, 3, X}, {-2, 4, X}, {-2, 5, X},
		{-1, -2, X}, {-1, -1, X}, {-1, 0, X}, {-1, 1, X}, {-1, 2, X}, {-1, 3, X}, {-1, 4, X}, {-1, 5, X},
		{0, -2, X}, {0, -1, X}, {0, 0, 1}, {0, 1, X}, {0, 2, X}, {0, 3, X}, {0, 4, X}, {0, 5, X},
		{1, -2, X}, {1, -1, X}, {1, 0, 1}, {1, 1, 1}, {1, 2, X}, {1, 3, X}, {1, 4, X}, {1, 5, X},
		{2, -2, X}, {2, -1, X}, {2, 0, 1}, {2, 1, 2}, {2, 2, 1}, {2, 3, X}, {2, 4, X}, {2, 5, X},
		{3, -2, X}, {3, -1, X}, {3, 0, 1}, {3, 1, 3}, {3, 2, 3}, {3, 3, 1}, {3, 4, X}, {3, 5, X},
		{4, -2, X}, {4, -1, X}, {4, 0, 1}, {4, 1, 4}, {4, 2, 6}, {4, 3, 4}, {4, 4, 1}, {4, 5, X},
		{5, -2, X}, {5, -1, X}, {5, 0, 1}, {5, 1, 5}, {5, 2, 10}, {5, 3, 10}, {5, 4, 5}, {5, 5, 1},
		{6, -2, X}, {6, -1, X}, {6, 0, 1}, {6, 1, 6}, {6, 2, 15}, {6, 3, 20}, {6, 4, 15}, {6, 5, 6},
		{7, -2, X}, {7, -1, X}, {7, 0, 1}, {7, 1, 7}, {7, 2, 21}, {7, 3, 35}, {7, 4, 35}, {7, 5, 21},
		{8, -2, X}, {8, -1, X}, {8, 0, 1}, {8, 1, 8}, {8, 2, 28}, {8, 3, 56}, {8, 4, 70}, {8, 5, 56},
		{9, -2, X}, {9, -1, X}, {9, 0, 1}, {9, 1, 9}, {9, 2, 36}, {9, 3, 84}, {9, 4, 126}, {9, 5, 126},
	}
	for _, c := range cases {
		got := combination.CombinationCount(c.n, c.r)
		if got != c.expected {
			t.Errorf("CombinationCount(%d, %d) = %d, expected: %d", c.n, c.r, got, c.expected)
		}
	}
}

func testCombination(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = i + 1
		}
		for r := -2; r <= 4; r++ {
			expected := Combination0(a, r)
			got := combination(a, r)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Combination(%v, %d) = %v, expected: %v", a, r, got, expected)
			}
		}
	}
}

func TestCombination1(t *testing.T) {
	testCombination(t, combination.Combination1[int])
}

func TestCombination2(t *testing.T) {
	testCombination(t, combination.Combination2[int])
}

func TestCombination2a(t *testing.T) {
	testCombination(t, combination.Combination2a[int])
}

func TestCombinationWithRepetitionCount(t *testing.T) {
	cases := []struct {
		n        int
		r        int
		expected int
	}{
		{-2, -2, X}, {-2, -1, X}, {-2, 0, X}, {-2, 1, X}, {-2, 2, X}, {-2, 3, X}, {-2, 4, X}, {-2, 5, X},
		{-1, -2, X}, {-1, -1, X}, {-1, 0, X}, {-1, 1, X}, {-1, 2, X}, {-1, 3, X}, {-1, 4, X}, {-1, 5, X},
		{0, -2, X}, {0, -1, X}, {0, 0, X}, {0, 1, X}, {0, 2, X}, {0, 3, X}, {0, 4, X}, {0, 5, X},
		{1, -2, X}, {1, -1, X}, {1, 0, 1}, {1, 1, 1}, {1, 2, 1}, {1, 3, 1}, {1, 4, 1}, {1, 5, 1},
		{2, -2, X}, {2, -1, X}, {2, 0, 1}, {2, 1, 2}, {2, 2, 3}, {2, 3, 4}, {2, 4, 5}, {2, 5, 6},
		{3, -2, X}, {3, -1, X}, {3, 0, 1}, {3, 1, 3}, {3, 2, 6}, {3, 3, 10}, {3, 4, 15}, {3, 5, 21},
		{4, -2, X}, {4, -1, X}, {4, 0, 1}, {4, 1, 4}, {4, 2, 10}, {4, 3, 20}, {4, 4, 35}, {4, 5, 56},
		{5, -2, X}, {5, -1, X}, {5, 0, 1}, {5, 1, 5}, {5, 2, 15}, {5, 3, 35}, {5, 4, 70}, {5, 5, 126},
		{6, -2, X}, {6, -1, X}, {6, 0, 1}, {6, 1, 6}, {6, 2, 21}, {6, 3, 56}, {6, 4, 126}, {6, 5, 252},
		{7, -2, X}, {7, -1, X}, {7, 0, 1}, {7, 1, 7}, {7, 2, 28}, {7, 3, 84}, {7, 4, 210}, {7, 5, 462},
		{8, -2, X}, {8, -1, X}, {8, 0, 1}, {8, 1, 8}, {8, 2, 36}, {8, 3, 120}, {8, 4, 330}, {8, 5, 792},
		{9, -2, X}, {9, -1, X}, {9, 0, 1}, {9, 1, 9}, {9, 2, 45}, {9, 3, 165}, {9, 4, 495}, {9, 5, 1287},
	}
	for _, c := range cases {
		got := combination.CombinationWithRepetitionCount(c.n, c.r)
		if got != c.expected {
			t.Errorf("CombinationWithRepetitionCount(%d, %d) = %d, expected: %d", c.n, c.r, got, c.expected)
		}
	}
}

func testCombinationWithRepetition(t *testing.T, combinationWithRepetition func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = i + 1
		}
		for r := -2; r <= 4; r++ {
			expected := CombinationWithRepetition0(a, r)
			got := combinationWithRepetition(a, r)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("CombinationWithRepetition(%v, %d) = %v, expected: %v", a, r, got, expected)
			}
		}
	}
}

func TestCombinationWithRepetition1(t *testing.T) {
	testCombinationWithRepetition(t, combination.CombinationWithRepetition1[int])
}

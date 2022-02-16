package combination_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/combination"
)

func TestCombinationCount(t *testing.T) {
	cases := []struct {
		n        int
		r        int
		expected int
	}{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 1},
		{2, 3, 0},
		{3, 0, 1},
		{3, 1, 3},
		{3, 2, 3},
		{3, 3, 1},
		{3, 4, 0},
	}
	for _, c := range cases {
		got := combination.CombinationCount(c.n, c.r)
		if got != c.expected {
			t.Errorf("CombinationCount(%d, %d) = %d, expected: %d", c.n, c.r, got, c.expected)
		}
	}
}

func testCombinationNeq0(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		expected := [][]int{{}}
		got := combination(s, 0)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Combination(%v, 0) = %v, expected: %v", s, got, expected)
		}
	}
}

func testCombinationNeq1(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		expected := [][]int{}
		for i := 0; i < n; i++ {
			expected = append(expected, []int{i})
		}
		got := combination(s, 1)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Combination(%v, 1) = %v, expected: %v", s, got, expected)
		}
	}
}

func testCombinationNeq2(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		expected := [][]int{}
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				expected = append(expected, []int{i, j})
			}
		}
		got := combination(s, 2)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Combination(%v, 2) = %v, expected: %v", s, got, expected)
		}
	}
}

func testCombinationNeq3(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		expected := [][]int{}
		for i := 0; i < n-2; i++ {
			for j := i + 1; j < n-1; j++ {
				for k := j + 1; k < n; k++ {
					expected = append(expected, []int{i, j, k})
				}
			}
		}
		got := combination(s, 3)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Combination(%v, 3) = %v, expected: %v", s, got, expected)
		}
	}
}

func testCombinationNeq4(t *testing.T, combination func([]int, int) [][]int) {
	for n := 0; n < 10; n++ {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		expected := [][]int{}
		for i := 0; i < n-3; i++ {
			for j := i + 1; j < n-2; j++ {
				for k := j + 1; k < n-1; k++ {
					for l := k + 1; l < n; l++ {
						expected = append(expected, []int{i, j, k, l})
					}
				}
			}
		}
		got := combination(s, 4)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Combination(%v, 4) = %v, expected: %v", s, got, expected)
		}
	}
}

func TestCombination1(t *testing.T) {
	testCombinationNeq0(t, combination.Combination1[int])
	testCombinationNeq1(t, combination.Combination1[int])
	testCombinationNeq2(t, combination.Combination1[int])
	testCombinationNeq3(t, combination.Combination1[int])
	testCombinationNeq4(t, combination.Combination1[int])
}

func TestCombination2(t *testing.T) {
	testCombinationNeq0(t, combination.Combination2[int])
	testCombinationNeq1(t, combination.Combination2[int])
	testCombinationNeq2(t, combination.Combination2[int])
	testCombinationNeq3(t, combination.Combination2[int])
	testCombinationNeq4(t, combination.Combination2[int])
}

func TestCombination2a(t *testing.T) {
	testCombinationNeq0(t, combination.Combination2a[int])
	testCombinationNeq1(t, combination.Combination2a[int])
	testCombinationNeq2(t, combination.Combination2a[int])
	testCombinationNeq3(t, combination.Combination2a[int])
	testCombinationNeq4(t, combination.Combination2a[int])
}

// func TestCombination3(t *testing.T) {
// 	testCombinationNeq0(t, combination.Combination3[int])
// 	testCombinationNeq1(t, combination.Combination3[int])
// 	testCombinationNeq2(t, combination.Combination3[int])
// 	testCombinationNeq3(t, combination.Combination3[int])
// 	testCombinationNeq4(t, combination.Combination3[int])
// }

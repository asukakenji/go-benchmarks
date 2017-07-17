package math_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common"
	"github.com/asukakenji/go-benchmarks/math"
)

func TestIsOdd(t *testing.T) {
	implementations := []struct {
		name string
		f    func(int) bool
	}{
		{"IsOddNaive", math.IsOddNaive},
		{"IsOdd0", math.IsOdd0},
		{"IsOdd1", math.IsOdd1},
	}
	cases := []struct {
		x        int
		expected bool
	}{
		{common.MinInt, false},
		{-3, true},
		{-2, false},
		{-1, true},
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{common.MaxInt, true},
	}
	for _, impl := range implementations {
		for _, c := range cases {
			got := impl.f(c.x)
			if got != c.expected {
				t.Errorf("%s(%d) = %t, expected %t", impl.name, c.x, got, c.expected)
			}
		}
	}
}

func TestIsEven(t *testing.T) {
	implementations := []struct {
		name string
		f    func(int) bool
	}{
		{"IsEvenNaive", math.IsEvenNaive},
		{"IsEven0", math.IsEven0},
		{"IsEven1", math.IsEven1},
	}
	cases := []struct {
		x        int
		expected bool
	}{
		{common.MinInt, true},
		{-3, false},
		{-2, true},
		{-1, false},
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{common.MaxInt, false},
	}
	for _, impl := range implementations {
		for _, c := range cases {
			got := impl.f(c.x)
			if got != c.expected {
				t.Errorf("%s(%d) = %t, expected %t", impl.name, c.x, got, c.expected)
			}
		}
	}
}

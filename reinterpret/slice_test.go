package reinterpret_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/reinterpret"
)

func TestIntSliceAsByteSlice(t *testing.T) {
	x := []int{0x0123456789abcdef}
	expected := []byte{0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01}
	got := reinterpret.IntSliceAsByteSlice(x)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IntSliceAsByteSlice(0x%x) = %v, expected %v", x[0], got, expected)
	}

	got[0] = 0x9a
	got[1] = 0xbc
	got[2] = 0xde
	got[3] = 0xf0
	got[4] = 0x12
	got[5] = 0x34
	got[6] = 0x56
	got[7] = 0x78
	expected2 := 0x78563412f0debc9a
	got2 := x[0]
	if got2 != expected2 {
		t.Errorf("IntSliceAsByteSlice(0x%x) = %v, expected %v", x[0], got, expected)
	}
}

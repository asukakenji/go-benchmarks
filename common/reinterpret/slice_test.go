package reinterpret_test

import (
	"reflect"
	"testing"

	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

func TestIntSliceAsByteSlice(t *testing.T) {
	x := []int{0x0123456789abcdef}
	expected := []byte{0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01}
	got := reinterpret.IntSliceAsByteSlice(x)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("IntSliceAsByteSlice([]int{0x%x}) = %v, expected %v", x[0], got, expected)
	}

	got[0] = 0x9a
	got[1] = 0xbc
	got[2] = 0xde
	got[3] = 0xf0
	got[4] = 0x12
	got[5] = 0x34
	got[6] = 0x56
	got[7] = 0x78
	expected2 := int(0x78563412f0debc9a)
	got2 := x[0]
	if got2 != expected2 {
		t.Errorf("IntSliceAsByteSlice([]int{0x%x}) = %v, expected %v", x[0], got2, expected2)
	}
}

func TestUintSliceAsByteSlice(t *testing.T) {
	x := []uint{0x0123456789abcdef}
	expected := []byte{0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01}
	got := reinterpret.UintSliceAsByteSlice(x)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("UintSliceAsByteSlice([]uint{0x%x}) = %v, expected %v", x[0], got, expected)
	}

	got[0] = 0xab
	got[1] = 0xcd
	got[2] = 0xef
	got[3] = 0x01
	got[4] = 0x23
	got[5] = 0x45
	got[6] = 0x67
	got[7] = 0x89
	expected2 := uint(0x8967452301efcdab)
	got2 := x[0]
	if got2 != expected2 {
		t.Errorf("UintSliceAsByteSlice([]uint{0x%x}) = %v, expected %v", x[0], got2, expected2)
	}
}

func TestUint32SliceAsByteSlice(t *testing.T) {
	x := []uint32{0x01234567}
	expected := []byte{0x67, 0x45, 0x23, 0x01}
	got := reinterpret.Uint32SliceAsByteSlice(x)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Uint32SliceAsByteSlice(uint32{0x%x}) = %v, expected %v", x[0], got, expected)
	}

	got[0] = 0x23
	got[1] = 0x45
	got[2] = 0x67
	got[3] = 0x89
	expected2 := uint32(0x89674523)
	got2 := x[0]
	if got2 != expected2 {
		t.Errorf("Uint32SliceAsByteSlice(uint32{0x%x}) = %v, expected %v", x[0], got2, expected2)
	}
}

func TestUint64SliceAsByteSlice(t *testing.T) {
	x := []uint64{0x0123456789abcdef}
	expected := []byte{0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01}
	got := reinterpret.Uint64SliceAsByteSlice(x)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Uint64SliceAsByteSlice(uint64{0x%x}) = %v, expected %v", x[0], got, expected)
	}

	got[0] = 0xab
	got[1] = 0xcd
	got[2] = 0xef
	got[3] = 0x01
	got[4] = 0x23
	got[5] = 0x45
	got[6] = 0x67
	got[7] = 0x89
	expected2 := uint64(0x8967452301efcdab)
	got2 := x[0]
	if got2 != expected2 {
		t.Errorf("Uint64SliceAsByteSlice(uint64{0x%x}) = %v, expected %v", x[0], got2, expected2)
	}
}

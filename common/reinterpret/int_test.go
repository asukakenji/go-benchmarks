package reinterpret_test

import (
	"testing"

	"github.com/asukakenji/go-benchmarks/common/reinterpret"
)

func TestIntAsUint(t *testing.T) {
	x := int(-8690466094656961759)
	expected := uint(0x8765432187654321)
	got := reinterpret.IntAsUint(x)
	if got != expected {
		t.Errorf("IntAsUint(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

func TestInt32AsUint32(t *testing.T) {
	x := int32(-2023406815)
	expected := uint32(0x87654321)
	got := reinterpret.Int32AsUint32(x)
	if got != expected {
		t.Errorf("Int32AsUint32(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

func TestInt64AsUint64(t *testing.T) {
	x := int64(-8690466094656961759)
	expected := uint64(0x8765432187654321)
	got := reinterpret.Int64AsUint64(x)
	if got != expected {
		t.Errorf("Int64AsUint64(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

func TestUintAsInt(t *testing.T) {
	x := uint(0x8765432187654321)
	expected := int(-8690466094656961759)
	got := reinterpret.UintAsInt(x)
	if got != expected {
		t.Errorf("UintAsInt(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

func TestUint32AsInt32(t *testing.T) {
	x := uint32(0x87654321)
	expected := int32(-2023406815)
	got := reinterpret.Uint32AsInt32(x)
	if got != expected {
		t.Errorf("Uint32AsInt32(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

func TestUint64AsInt64(t *testing.T) {
	x := uint64(0x8765432187654321)
	expected := int64(-8690466094656961759)
	got := reinterpret.Uint64AsInt64(x)
	if got != expected {
		t.Errorf("Uint64AsInt64(0x%x) = 0x%x, expected 0x%x", x, got, expected)
	}
}

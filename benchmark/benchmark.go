package benchmark

import "testing"

/*

// From "reflect":

const (
	Invalid       Kind = iota // 0
	Bool                      // 1
	Int                       // 2
	Int8                      // 3
	Int16                     // 4
	Int32                     // 5
	Int64                     // 6
	Uint                      // 7
	Uint8                     // 8
	Uint16                    // 9
	Uint32                    // 10
	Uint64                    // 11
	Uintptr                   // 12
	Float32                   // 13
	Float64                   // 14
	Complex64                 // 15
	Complex128                // 16
	Array                     // 17
	Chan                      // 18
	Func                      // 19
	Interface                 // 20
	Map                       // 21
	Ptr                       // 22
	Slice                     // 23
	String                    // 24
	Struct                    // 25
	UnsafePointer             // 26
)
*/

// UintFuncUint benchmarks a function with the signature:
//     func(uint) uint
// ID: B-7-7
func UintFuncUint(b *testing.B, genUintFunc func() uint, f func(uint) uint) {
	for i, count := 0, b.N; i < count; i++ {
		f(genUintFunc())
	}
}

// UintFuncUint32 benchmarks a function with the signature:
//     func(uint32) uint
// ID: B-7-10
func UintFuncUint32(b *testing.B, genUint32Func func() uint32, f func(uint32) uint) {
	for i, count := 0, b.N; i < count; i++ {
		f(genUint32Func())
	}
}

// UintFuncUint64 benchmarks a function with the signature:
//     func(uint64) uint
// ID: B-7-11
func UintFuncUint64(b *testing.B, genUint64Func func() uint64, f func(uint64) uint) {
	for i, count := 0, b.N; i < count; i++ {
		f(genUint64Func())
	}
}

package reinterpret

import (
	"reflect"
	"unsafe"

	"github.com/asukakenji/go-benchmarks/common"
)

// IntSliceAsByteSlice reinterprets a []int as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-2-8 (byte = uint8)
func IntSliceAsByteSlice(x []int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) * common.SizeOfIntInBytes,
		Cap:  cap(x) * common.SizeOfIntInBytes,
	}))
}

// UintSliceAsByteSlice reinterprets a []uint as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-7-8 (byte = uint8)
func UintSliceAsByteSlice(x []uint) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) * common.SizeOfIntInBytes,
		Cap:  cap(x) * common.SizeOfIntInBytes,
	}))
}

// Uint32SliceAsByteSlice reinterprets a []uint32 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-10-8 (byte = uint8)
func Uint32SliceAsByteSlice(x []uint32) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 2,
		Cap:  cap(x) << 2,
	}))
}

// Uint64SliceAsByteSlice reinterprets a []uint64 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-11-8 (byte = uint8)
func Uint64SliceAsByteSlice(x []uint64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 3,
		Cap:  cap(x) << 3,
	}))
}

package reinterpret

import (
	"reflect"
	"unsafe"

	"github.com/asukakenji/go-benchmarks/common"
)

// IntSliceAsByteSlice reinterprets a []int as a []byte.
// Since the returned slice does not contribute in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
// ID: RS-2-8 (byte = uint8)
func IntSliceAsByteSlice(x []int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) * common.SizeOfIntInBytes,
		Cap:  cap(x) * common.SizeOfIntInBytes,
	}))
}

// UintSliceAsByteSlice reinterprets a []uint as a []byte.
// Since the returned slice does not contribute in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
// ID: RS-7-8 (byte = uint8)
func UintSliceAsByteSlice(x []uint) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) * common.SizeOfIntInBytes,
		Cap:  cap(x) * common.SizeOfIntInBytes,
	}))
}

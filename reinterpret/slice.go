package reinterpret

import (
	"reflect"
	"unsafe"

	"github.com/asukakenji/go-benchmarks/common"
)

// IntSliceAsByteSlice reinterprets an []int as []byte.
// Since the returned slice does not contribute in reference counting,
// the caller is responsible to keep a proper reference to the actual data
// to prevent data lost by garbage collection.
// A variable of type []int referencing to the same slice as x is an example of
// a proper reference.
// ID: RS-2-8 (byte = uint8)
func IntSliceAsByteSlice(x []int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) * common.SizeOfIntInBytes,
		Cap:  cap(x) * common.SizeOfIntInBytes,
	}))
}

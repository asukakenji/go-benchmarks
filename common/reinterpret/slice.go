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

// Int8SliceAsByteSlice reinterprets a []int8 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-3-8 (byte = uint8)
func Int8SliceAsByteSlice(x []int8) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x),
		Cap:  cap(x),
	}))
}

// Int16SliceAsByteSlice reinterprets a []int16 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-4-8 (byte = uint8)
func Int16SliceAsByteSlice(x []int16) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 1,
		Cap:  cap(x) << 1,
	}))
}

// Int32SliceAsByteSlice reinterprets a []int32 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-5-8 (byte = uint8)
func Int32SliceAsByteSlice(x []int32) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 2,
		Cap:  cap(x) << 2,
	}))
}

// Int64SliceAsByteSlice reinterprets a []int64 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-6-8 (byte = uint8)
func Int64SliceAsByteSlice(x []int64) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 3,
		Cap:  cap(x) << 3,
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

// Uint8SliceAsByteSlice reinterprets a []uint8 as a []byte.
// Since byte is an alias of uint8, x is returned.
//
// ID: RS-8-8 (byte = uint8)
func Uint8SliceAsByteSlice(x []uint8) []byte {
	return x
}

// Uint16SliceAsByteSlice reinterprets a []uint16 as a []byte.
// Since the returned slice does not participate in reference counting,
// the caller is responsible to keep a reference to the original slice
// to prevent it from being garbage collected.
//
// ID: RS-9-8 (byte = uint8)
func Uint16SliceAsByteSlice(x []uint16) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.SliceHeader)(unsafe.Pointer(&x)).Data,
		Len:  len(x) << 2,
		Cap:  cap(x) << 2,
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

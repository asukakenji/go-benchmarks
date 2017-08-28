package reinterpret

import "unsafe"

// IntAsUint reinterprets an int value bit-by-bit as an uint value.
// ID: R-2-7
func IntAsUint(x int) uint {
	return *(*uint)(unsafe.Pointer(&x))
}

// Int8AsUint8 reinterprets an int8 value bit-by-bit as an uint8 value.
// ID: R-3-8
func Int8AsUint8(x int8) uint8 {
	return *(*uint8)(unsafe.Pointer(&x))
}

// Int16AsUint16 reinterprets an int16 value bit-by-bit as an uint16 value.
// ID: R-4-9
func Int16AsUint16(x int16) uint16 {
	return *(*uint16)(unsafe.Pointer(&x))
}

// Int32AsUint32 reinterprets an int32 value bit-by-bit as an uint32 value.
// ID: R-5-10
func Int32AsUint32(x int32) uint32 {
	return *(*uint32)(unsafe.Pointer(&x))
}

// Int64AsUint64 reinterprets an int64 value bit-by-bit as an uint64 value.
// ID: R-6-11
func Int64AsUint64(x int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&x))
}

// UintAsInt reinterprets an uint value bit-by-bit as an int value.
// ID: R-7-2
func UintAsInt(x uint) int {
	return *(*int)(unsafe.Pointer(&x))
}

// Uint8AsInt8 reinterprets an uint8 value bit-by-bit as an int8 value.
// ID: R-8-3
func Uint8AsInt8(x uint8) int8 {
	return *(*int8)(unsafe.Pointer(&x))
}

// Uint16AsInt16 reinterprets an uint16 value bit-by-bit as an int16 value.
// ID: R-9-4
func Uint16AsInt16(x uint16) int16 {
	return *(*int16)(unsafe.Pointer(&x))
}

// Uint32AsInt32 reinterprets an uint32 value bit-by-bit as an int32 value.
// ID: R-10-5
func Uint32AsInt32(x uint32) int32 {
	return *(*int32)(unsafe.Pointer(&x))
}

// Uint64AsInt64 reinterprets an uint64 value bit-by-bit as an int64 value.
// ID: R-11-6
func Uint64AsInt64(x uint64) int64 {
	return *(*int64)(unsafe.Pointer(&x))
}

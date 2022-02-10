package benchmarks

import "constraints"

/////////////////
// Size of int //
/////////////////

const (
	// On 32-bit platforms: 0xFFFFFFFF
	// On 64-bit platforms: 0xFFFFFFFFFFFFFFFF
	MaxUint = ^uint(0)
	// On 32-bit platforms: 2
	// On 64-bit platforms: 3
	LogSizeOfUintInBytes = MaxUint>>8&1 + MaxUint>>16&1 + MaxUint>>32&1
	// On 32-bit platforms: 4
	// On 64-bit platforms: 8
	SizeOfUintInBytes = 1 << LogSizeOfUintInBytes
	// On 32-bit platforms: 32
	// On 64-bit platforms: 64
	SizeOfUintInBits = SizeOfUintInBytes << 3
)

func SizeOf[T constraints.Unsigned]() int {
	//  8-bit: 0xFF
	// 16-bit: 0xFFFF
	// 32-bit: 0xFFFFFFFF
	// 64-bit: 0xFFFFFFFFFFFFFFFF
	maxT := ^T(0)
	//  8-bit: 0
	// 16-bit: 1
	// 32-bit: 2
	// 64-bit: 3
	logSizeOfTInBytes := maxT>>8&1 + maxT>>16&1 + maxT>>32&1
	//  8-bit: 1
	// 16-bit: 2
	// 32-bit: 4
	// 64-bit: 8
	sizeOfTInBytes := 1 << logSizeOfTInBytes
	//  8-bit: 8
	// 16-bit: 16
	// 32-bit: 32
	// 64-bit: 64
	sizeOfTInBits := sizeOfTInBytes << 3
	return sizeOfTInBits
}

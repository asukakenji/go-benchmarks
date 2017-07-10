package common

// Integer limit values.
const (
	MaxUint             = ^uint(0)
	MaxInt              = int(MaxUint >> 1)
	MinInt              = ^MaxInt
	LogSizeOfIntInBytes = MaxUint>>8&1 + MaxUint>>16&1 + MaxUint>>32&1
	SizeOfIntInBytes    = 1 << LogSizeOfIntInBytes
	SizeOfIntInBits     = SizeOfIntInBytes << 3
)

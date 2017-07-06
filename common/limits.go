package common

const (
	MaxUint              = ^uint(0)
	MaxInt               = int(MaxUint >> 1)
	MinInt               = ^MaxInt
	LogSizeOfUintInBytes = MaxUint>>8&1 + MaxUint>>16&1 + MaxUint>>32&1
	SizeOfUintInBytes    = 1 << LogSizeOfUintInBytes
	SizeOfUintInBits     = SizeOfUintInBytes << 3
)

package common

import "os"

// Page size
var (
	PageSizeInBytes uint
	IntCountPerPage uint
)

func init() {
	PageSizeInBytes = uint(os.Getpagesize())
	IntCountPerPage = PageSizeInBytes / SizeOfIntInBytes
}

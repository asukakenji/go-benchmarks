// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5a)
package subtract

// NOTE: The original code compiles to the same assembly on darwin_amd64.
// func OriginalCode(x uint) int {
// 	sum := int(x)
// 	for x != 0 {
// 		x = x >> 1
// 		sum = sum - int(x)
// 	}
// 	return sum
// }

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	sum := int(x)
	for x != 0 {
		x >>= 1
		sum -= int(x)
	}
	return sum
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	sum := int8(x)
	for x != 0 {
		x >>= 1
		sum -= int8(x)
	}
	return int(sum)
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	sum := int16(x)
	for x != 0 {
		x >>= 1
		sum -= int16(x)
	}
	return int(sum)
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	sum := int32(x)
	for x != 0 {
		x >>= 1
		sum -= int32(x)
	}
	return int(sum)
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	sum := int64(x)
	for x != 0 {
		x >>= 1
		sum -= int64(x)
	}
	return int(sum)
}

// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop4)
package reset

// NOTE: The original code compiles to the same assembly on darwin_amd64.
// func OriginalCode(x uint) int {
// 	n := 0
// 	for x != 0 {
// 		n = n + 1
// 		x = x & (x - 1)
// 	}
// 	return n
// }

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	n := 0
	for x != 0 {
		n++
		x &= (x - 1)
	}
	return n
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	n := 0
	for x != 0 {
		n++
		x &= (x - 1)
	}
	return n
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	n := 0
	for x != 0 {
		n++
		x &= (x - 1)
	}
	return n
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	n := 0
	for x != 0 {
		n++
		x &= (x - 1)
	}
	return n
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	n := 0
	for x != 0 {
		n++
		x &= (x - 1)
	}
	return n
}

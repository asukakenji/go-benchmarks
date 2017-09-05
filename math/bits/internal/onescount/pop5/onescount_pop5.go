// Source: http://www.hackersdelight.org/hdcodetxt/pop.c.txt (pop5)
package pop5

import "github.com/asukakenji/go-benchmarks/common/reinterpret"

func rotate32(x uint32, n uint) uint32 {
	if n > 63 {
		panic("rotate32, n out of range.")
	}
	return (x << n) | (x >> (32 - n))
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	sum := reinterpret.Uint32AsInt32(x)
	for i := 1; i <= 31; i++ {
		x = rotate32(x, 1)
		sum = sum + reinterpret.Uint32AsInt32(x)
	}
	return int(-sum)
}

func rotate64(x uint64, n uint) uint64 {
	if n > 127 {
		panic("rotate64, n out of range.")
	}
	return (x << n) | (x >> (64 - n))
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	sum := reinterpret.Uint64AsInt64(x)
	for i := 1; i <= 63; i++ {
		x = rotate64(x, 1)
		sum = sum + reinterpret.Uint64AsInt64(x)
	}
	return int(-sum)
}

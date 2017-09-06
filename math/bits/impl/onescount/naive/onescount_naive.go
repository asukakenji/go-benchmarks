package naive

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int {
	count := 0
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int {
	count := int(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int {
	count := int(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	count := int(0)
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

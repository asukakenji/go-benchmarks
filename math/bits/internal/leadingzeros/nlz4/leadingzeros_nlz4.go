package nlz4

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int {
	var y, m, n int32

	y = -int32(x >> 16) // If left half of x is 0,
	m = (y >> 16) & 16  // set n = 16.  If left half
	n = 16 - m          // is nonzero, set n = 0 and
	x = x >> uint32(m)  // shift x right 16.
	// Now x is of the form 0000xxxx.
	y = int32(x - 0x100) // If positions 8-15 are 0,
	m = (y >> 16) & 8    // add 8 to n and shift x left 8.
	n = n + m
	x = x << uint32(m)

	y = int32(x - 0x1000) // If positions 12-15 are 0,
	m = (y >> 16) & 4     // add 4 to n and shift x left 4.
	n = n + m
	x = x << uint32(m)

	y = int32(x - 0x4000) // If positions 14-15 are 0,
	m = (y >> 16) & 2     // add 2 to n and shift x left 2.
	n = n + m
	x = x << uint32(m)

	y = int32(x >> 14) // Set y = 0, 1, 2, or 3.
	m = y &^ (y >> 1)  // Set m = 0, 1, 2, or 2 resp.
	return int(n + 2 - m)
}

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
// func LeadingZeros64(x uint64) int {
// 	var y, m, n int64
//
// 	y = -int64(x >> 32) // If left half of x is 0,
// 	m = (y >> 32) & 32  // set n = 32.  If left half
// 	n = 32 - m          // is nonzero, set n = 0 and
// 	x = x >> uint64(m)  // shift x right 32.
// 	// Now x is of the form 00000000xxxxxxxx.
// 	y = int64(x - 0x10000) // If positions 16-31 are 0,
// 	m = (y >> 32) & 16     // add 16 to n and shift x left 16.
// 	n = n + m
// 	x = x << uint64(m)
//
// 	y = int64(x - 0x100000) // If positions 20-31 are 0,
// 	m = (y >> 32) & 8       // add 8 to n and shift x left 8.
// 	n = n + m
// 	x = x << uint64(m)
//
// 	y = int64(x - 0x1000000) // If positions 24-31 are 0,
// 	m = (y >> 32) & 4        // add 4 to n and shift x left 4.
// 	n = n + m
// 	x = x << uint64(m)
//
// 	y = int64(x - 0x4000000) // If positions 26-31 are 0,
// 	m = (y >> 32) & 2        // add 2 to n and shift x left 2.
// 	n = n + m
// 	x = x << uint64(m)
//
// 	y = int64(x >> 30) // Set y = 0, 1, 2, or 3.
// 	m = y &^ (y >> 1)  // Set m = 0, 1, 2, or 2 resp.
// 	return int(n + 2 - m)
// }

# Math

## BitCount

### 32-bit

The most efficient implementations for 32-bit are:

`BitCountUint32Pop1Alt`:

```go
func BitCountUint32Pop1Alt(x uint32) uint {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	return uint((x * 0x01010101) >> 24)
}
```

`BitCountUint32Pop3`:

```go
func BitCountUint32Pop3(x uint32) uint {
	n := (x >> 1) & 0x77777777 // Count bits in
	x = x - n                  // each 4-bit
	n = (n >> 1) & 0x77777777  // field.
	x = x - n
	n = (n >> 1) & 0x77777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f // Get byte sums.
	x = x * 0x01010101              // Add the bytes.
	return uint(x >> 24)
}
```

`BitCountUint32Pop6`:

```go
var byteToBitCountTable = [...]uint{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,

	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

func BitCountUint32Pop6(x uint32) uint {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)]
}
```

The implementations `BitCountUint32Pop1Alt` and `BitCountUint32Pop6` are more widely used. They are easier than `BitCountUint32Pop3` to understand. As a result, they are used as the [optimized implementation](https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L847-L850) and [fallback implementation](https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L852-L857) of GCC's `__builtin_popcount` respectively. In benchmarks, these 3 implementations score almost the same most of the time. `BitCountUint32Pop1Alt` is usually the best among them when there is a difference.

### 64-bit

Here are the 64-bit implementations:

`BitCountUint64Pop1Alt`:

```go
func BitCountUint64Pop1Alt(x uint64) uint {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	return uint((x * 0x0101010101010101) >> 56)
}
```

`BitCountUint64Pop3`:

```go
func BitCountUint64Pop3(x uint64) uint {
	n := (x >> 1) & 0x7777777777777777 // Count bits in
	x = x - n                          // each 4-bit
	n = (n >> 1) & 0x7777777777777777  // field.
	x = x - n
	n = (n >> 1) & 0x7777777777777777
	x = x - n
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f // Get byte sums.
	x = x * 0x0101010101010101              // Add the bytes.
	return uint(x >> 56)
}
```

`BitCountUint64Pop6`:

```go
// byteToBitCountTable is the same as that in the 32-bit implementation above.

func BitCountUint64Pop6(x uint64) uint {
	return byteToBitCountTable[x&0xff] +
		byteToBitCountTable[(x>>8)&0xff] +
		byteToBitCountTable[(x>>16)&0xff] +
		byteToBitCountTable[(x>>24)&0xff] +
		byteToBitCountTable[(x>>32)&0xff] +
		byteToBitCountTable[(x>>40)&0xff] +
		byteToBitCountTable[(x>>48)&0xff] +
		byteToBitCountTable[(x>>56)]
}
```

For 64-bit, the best implementations are `BitCountUint64Pop1Alt` and `BitCountUint64Pop3`. The table-based implementation `BitCountUint64Pop6` is significantly slower than its 32-bit counterpart.

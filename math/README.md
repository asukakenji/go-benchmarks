# Math

## BitCount

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
var table = [...]uint{
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
	return table[x&0xff] +
		table[(x>>8)&0xff] +
		table[(x>>16)&0xff] +
		table[(x>>24)]
}
```

The implementations `BitCountUint32Pop1Alt` and `BitCountUint32Pop6` are more widely used. They are easier than `BitCountUint32Pop3` to understand. As a result, they are used as the [optimized implementation](https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L847-L850) and [fallback implementation](https://github.com/gcc-mirror/gcc/blob/master/libgcc/libgcc2.c#L852-L857) of GCC's `__builtin_popcount` respectively.

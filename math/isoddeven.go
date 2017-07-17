package math

// IsOddNaive returns whether x is an odd number.
func IsOddNaive(x int) bool {
	return x%2 != 0
}

// IsOdd0 returns whether x is an odd number.
func IsOdd0(x int) bool {
	return x&1 != 0
}

// IsOdd1 returns whether x is an odd number.
func IsOdd1(x int) bool {
	return x&1 == 1
}

// IsEvenNaive returns whether x is an even number.
func IsEvenNaive(x int) bool {
	return x%2 == 0
}

// IsEven0 returns whether x is an even number.
func IsEven0(x int) bool {
	return x&1 == 0
}

// IsEven1 returns whether x is an even number.
func IsEven1(x int) bool {
	return x&1 != 1
}

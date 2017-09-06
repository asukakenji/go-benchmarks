package gccbuiltin

/*
#include <stdint.h>

// These functions calls the built-ins in GCC / clang.

static inline int popcount32(uint32_t x)
{
	if (sizeof(unsigned int) == sizeof(uint32_t))
	{
		return __builtin_popcount((unsigned int) x);
	}
	else if (sizeof(unsigned long) == sizeof(uint32_t))
	{
		return __builtin_popcountl((unsigned long) x);
	}
	else if (sizeof(unsigned long long) == sizeof(uint32_t))
	{
		return __builtin_popcountll((unsigned long long) x);
	}
	else
	{
		return *((int*) 0);
	}
}

static inline int popcount64(uint64_t x)
{
	if (sizeof(unsigned int) == sizeof(uint64_t))
	{
		return __builtin_popcount((unsigned int) x);
	}
	else if (sizeof(unsigned long) == sizeof(uint64_t))
	{
		return __builtin_popcountl((unsigned long) x);
	}
	else if (sizeof(unsigned long long) == sizeof(uint64_t))
	{
		return __builtin_popcountll((unsigned long long) x);
	}
	else
	{
		return *((int*) 0);
	}
}
*/
import "C"

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int {
	return int(C.popcount32(C.uint32_t(x)))
}

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	return int(C.popcount64(C.uint64_t(x)))
}

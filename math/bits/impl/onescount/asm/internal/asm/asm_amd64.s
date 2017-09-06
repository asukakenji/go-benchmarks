#include "textflag.h"

TEXT ·OnesCount64Asm(SB),NOSPLIT,$0
	MOVQ x+0(FP), AX
	POPCNTQ AX, AX
	MOVQ AX, ret+8(FP)
	RET

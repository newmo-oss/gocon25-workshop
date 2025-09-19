//go:build amd64

// Go アセンブリで関数を実装します
// 参考: https://go.dev/doc/asm

#include "textflag.h"

// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0-24
	MOVQ a+0(FP), AX    // 引数 a を AX レジスタに読み込む
	MOVQ b+8(FP), BX    // 引数 b を BX レジスタに読み込む
	ADDQ BX, AX         // AX = AX + BX
	MOVQ AX, ret+16(FP) // 結果を戻り値の位置に書き込む
	RET

// func Sub(a, b int64) int64
TEXT ·Sub(SB), NOSPLIT, $0-24
	MOVQ a+0(FP), AX    // 引数 a を AX レジスタに読み込む
	MOVQ b+8(FP), BX    // 引数 b を BX レジスタに読み込む
	SUBQ BX, AX         // AX = AX - BX (引き算)
	MOVQ AX, ret+16(FP) // 結果を戻り値の位置に書き込む
	RET
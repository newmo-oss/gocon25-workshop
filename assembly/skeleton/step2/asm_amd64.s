//go:build amd64

// Go アセンブリで関数を実装します
// 参考: https://go.dev/doc/asm

#include "textflag.h"

// ============================================
// 実装タスク
// ============================================

// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0-24
	// TODO: 引数 a を AX レジスタに読み込む
	
	// TODO: 引数 b を BX レジスタに読み込む
	
	// TODO: AX と BX を加算（結果は AX に格納される）
	
	// TODO: 結果を戻り値の位置に書き込む
	
	RET

// func Sub(a, b int64) int64
TEXT ·Sub(SB), NOSPLIT, $0-24
	// TODO: 引き算を実装
	
	RET

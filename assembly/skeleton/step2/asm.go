package main

// アセンブリで実装する関数の宣言
// これらの関数の実体は asm_amd64.s に記述します

// === 実装タスク ===

// Add は2つの整数を足し算します
//
//go:noescape
func Add(a, b int64) int64

// Sub は2つの整数を引き算します
//
//go:noescape
func Sub(a, b int64) int64

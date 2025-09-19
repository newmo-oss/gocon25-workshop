package main

import "testing"

// ====================================
// 実装タスクのテスト
// ====================================

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int64
	}{
		{1, 2, 3},
		{10, 20, 30},
		{-5, 5, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b, want int64
	}{
		{5, 3, 2},
		{10, 10, 0},
		{-5, -3, -2},
		{100, 50, 50},
	}

	for _, tt := range tests {
		got := Sub(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Sub(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

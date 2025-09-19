// sample.go - toolexec の動作確認用プログラム
package main

import (
	"fmt"
)

// ビルド時に -ldflags で値を設定できる変数
var (
	buildTime    = "unknown"
	buildVersion = "dev"
)

func main() {
	fmt.Println(sayFromGoConference())
}

func sayFromGoConference() string {
	return "Hello, Gopher Welcome to Go Conference 2025!"
}

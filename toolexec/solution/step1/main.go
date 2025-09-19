package main

import (
	"os"
	"os/exec"
)

func main() {
	// toolexec には最低2つの引数が渡されます
	// Args[0]: このプログラム自身のパス
	// Args[1]: 実行するツール（compile, link など）のパス
	// Args[2:]: ツールに渡す引数
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	toolPath := os.Args[1]
	toolArgs := os.Args[2:]

	// 元のツールを実行するコマンドを作成
	cmd := exec.Command(toolPath, toolArgs...)

	// 標準出力と標準エラー出力を設定
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// コマンドを実行
	err := cmd.Run()

	// エラーがあった場合、適切な終了コードで終了
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

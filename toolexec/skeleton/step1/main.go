package main

import (
	"os"
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

	// TODO: exec.Command を使って元のツールを実行するコマンドを作成してください
	// ヒント: exec.Command(toolPath, toolArgs...)

	// TODO: 標準出力と標準エラー出力を設定してください
	// ヒント: cmd.Stdout = os.Stdout
	// ヒント: cmd.Stderr = os.Stderr

	// TODO: コマンドを実行してください
	// ヒント: cmd.Run()

	// TODO: エラーがあった場合、適切な終了コードで終了してください
	// ヒント: exec.ExitError 型を使って終了コードを取得
}

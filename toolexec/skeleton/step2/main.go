package main

import (
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	toolPath := os.Args[1]
	toolArgs := os.Args[2:]

	// TODO: filepath.Base() を使ってツール名（compile, link など）を取得してください
	// ヒント: toolName := filepath.Base(toolPath)

	// TODO: os.Getenv() を使って TOOLEXEC_IMPORTPATH 環境変数を取得してください
	// この環境変数には現在ビルド中のパッケージ名が入っています
	// ヒント: importPath := os.Getenv("TOOLEXEC_IMPORTPATH")

	// TODO: ツール名とパッケージ名を標準エラー出力に表示してください
	// fmt.Fprintf(os.Stderr, "[TOOLEXEC] ...\n", ...)
	// 例: [TOOLEXEC] Running compile for package main

	// Step 1 のコードをここに（元のツールの実行）
	cmd := exec.Command(toolPath, toolArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

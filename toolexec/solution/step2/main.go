package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	toolPath := os.Args[1]
	toolArgs := os.Args[2:]

	// ツール名を取得（compile, link など）
	toolName := filepath.Base(toolPath)

	// TOOLEXEC_IMPORTPATH 環境変数を取得
	// この環境変数には現在ビルド中のパッケージ名が入っています
	importPath := os.Getenv("TOOLEXEC_IMPORTPATH")

	// ツール名とパッケージ名を表示
	if importPath != "" {
		fmt.Fprintf(os.Stderr, "[TOOLEXEC] Running %s for package %s\n", toolName, importPath)
	} else {
		fmt.Fprintf(os.Stderr, "[TOOLEXEC] Running %s\n", toolName)
	}

	// 元のツールの実行
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

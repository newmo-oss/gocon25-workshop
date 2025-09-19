package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Gopher のASCIIアート
const gopher = `
   D;;:;;:;;:;;:;:;:;;:;:;:;:;;:;;:;;:;;:;z
   $;:;:;::;::;:;;:;;:;;:;;:;;::;;::;;:;;;I
  .I;;:;;:;;:;;::;;:;:;:;;:;:;;:;:;:;::;;:I
  ,<;;::;:;;::;;:;;;;;;;;:;::;;:;;:;;;:;;;I
  ,(;;;:;::;:;;::;;j=1J71<;;;:;:;;::;:;:;:I
  J;;:;;;:;;::;;;;:r  ] .>;;;:;:;:;;:;:;;;r
  z;;::;:;;:;;:;;j=<?75?7~?I;;:;;:;;;:;:;<]
  (<;;;;;;:;;;;;;?+~(J-J-_(3;;;;;;::;;:;;+\
  ,(;:;:;j/7!''??1+?MMMMM1+?7771+<;;;:;;:j
  .P;;;;J!..       4;<<iJ        .4<;;:;;2 
.3;J<;;j\(M#Q       D;<2.MM5.      1:;;;j73,
$;jMN<;?|,WH3       $;:t.MM#       ,(;;jP;;?|
4<;T9TJ;?.        .J;;;?&         .t;;jM@:;+%
 (1++++Y+;?C+...J7<;;;:;;?i..  ..J>;jv<;;;j=
         .71+<;;;;;;;:;;;;;;;;;;<+J=  ?77!
             '_?771+++++++++?77!
`

var compileShown bool // compile時にGopherを表示したかどうかのフラグ

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	toolPath := os.Args[1]
	toolArgs := os.Args[2:]
	toolName := filepath.Base(toolPath)

	// ツールに応じてGopherとプログレスバーを表示
	switch toolName {
	case "compile":
		if !compileShown && os.Getenv("NO_GOPHER") != "1" {
			fmt.Fprintln(os.Stderr, "\n=== Go Build with Gopher ===")
			fmt.Fprint(os.Stderr, gopher)
			compileShown = true

			// TODO: コンパイル中のプログレスバーを表示
			// ヒント: showProgress("コンパイル中", 1*time.Second)
		}

	case "link":
		// TODO: リンク中のプログレスバーを表示
		// ヒント: showProgress("リンク中", 500*time.Millisecond)

		// TODO: ビルド完了メッセージを表示
		// ヒント: fmt.Fprintln(os.Stderr, "\n🎉 ビルド完了！")
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

// showProgress はプログレスバーを表示します（実装済み）
func showProgress(message string, duration time.Duration) {
	fmt.Fprintf(os.Stderr, "%s ", message)
	steps := 30
	for i := 0; i < steps; i++ {
		fmt.Fprint(os.Stderr, "█")
		time.Sleep(duration / time.Duration(steps))
	}
	fmt.Fprintln(os.Stderr, " ✅")
}

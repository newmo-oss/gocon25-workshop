author: sivchari
summary: Deep Dive into toolexec 
id: toolexec
categories: codelab,markdown
environments: Web
status: Published

# Go toolexec Codelab

このコードラボでは、Go の `-toolexec` フラグを使用してビルドプロセスをカスタマイズする方法を学習します。

## 進め方

1. 各ステップの「学習内容」と「ゴール」を確認してください
2. `skeleton/stepX/main.go` の TODO コメントを参考にして実装してください
3. 実装したら`testdata/sample.go`に実装したプログラムをtoolexecフラグで渡して動作を確認します

```bash
go run skeleton/stepX/main.go
go build -toolexec="$PWD/skeleton/stepX/main" testdata/sample.go
```

4. わからない場合は「ヒント」を参照してください

---

## Step 1: toolexec で動作するプログラムを実装する

### ゴール

元のツールをそのまま実行する toolexec プログラムが動作することを確認する

### 学習内容

`-toolexec` フラグを使うと、Go のビルドツール（compile, link など）の実行に任意のプログラムを適用することができます。例えば[DataDog/orchestrion](https://github.com/DataDog/orchestrion)や[alibaba/loongsuite-go-agent](https://github.com/alibaba/loongsuite-go-agent)ではビルド前に自動計装のためのコードを差し込むために使用しています。

最初のステップでは、単に元のツールを呼び出すだけの透過的なラッパーを作ります。

### toolexec の基本構造

```
go build -toolexec="mytool" main.go
    ↓
mytool /path/to/compile [compile の引数...]
    ↓
mytool /path/to/link [link の引数...]
```

**toolexec プログラムが受け取る引数：**
- `os.Args[0]`: toolexec プログラム自身のパス
- `os.Args[1]`: 実行するツール（compile, link など）のパス
- `os.Args[2:]`: ツールに渡す引数

### 実装タスク

`skeleton/step1/main.go` を開いて、TODO コメントの箇所を実装してください。

必要な実装：
1. `exec.Command` で元のツールを実行するコマンドを作成
2. 標準出力と標準エラー出力を設定
3. コマンドを実行
4. 終了コードを適切に処理

### 実行して確認

```bash
# Step 1 のプログラムをビルド
cd skeleton/step1
go build -o mytoolexec main.go

# toolexec として使用
go build -toolexec="$PWD/mytoolexec" ../../testdata/sample.go

# 正常にビルドできれば成功！
./sample
```

期待される出力：
```
=== Sample Program ===
Hello, Gopher Welcome to Go Conference 2025!
```

---

## Step 2: ビルドプロセスを観察する

### ゴール
ビルドで実行されるツール（compile、link など）を可視化する

### 学習内容

Go のビルドは複数のツールが連携して動作します：
- **compile**: .go ファイルを .o ファイルにコンパイル
- **link**: .o ファイルを実行ファイルにリンク

toolexec を使用して`go build`で実際にどのようなツールが内部で動いているか観察してみましょう。

### 観察したい情報

1. **どのツールが呼ばれたか**: compile? link? その他?
2. **何をビルドしているか**: main パッケージ？標準ライブラリ？
3. **実行順序**: どの順番でツールが動くか？

### 使える情報源

- `os.Args[1]`: 実行されるツールのパス（例: `/usr/local/go/pkg/tool/darwin_amd64/compile`）
- `TOOLEXEC_IMPORTPATH`: 環境変数。現在ビルド中のパッケージ名

### 実装タスク

`skeleton/step2/main.go` を開いて、以下を実装してください：

```go
// 1. ツール名を取得（パスから最後の部分だけ）
toolName := filepath.Base(toolPath)

// 2. ビルド対象のパッケージ名を環境変数から取得
pkg := os.Getenv("TOOLEXEC_IMPORTPATH")
if pkg == "" {
    pkg = "(unknown)"
}

// 3. 何が起きているか表示
fmt.Fprintf(os.Stderr, "[TOOLEXEC] Running %s for package %s\n", toolName, pkg)
```

### 実行して確認

```bash
# Step 2 のプログラムをビルド
cd skeleton/step2
go build -o mytoolexec main.go

# toolexec として使用（-a で全パッケージを再ビルド）
go build -a -toolexec="$PWD/mytoolexec" ../../testdata/sample.go 2>&1 | grep TOOLEXEC
```

期待される出力例：
```
# internal/unsafeheader
[TOOLEXEC] Running compile for package internal/unsafeheader
# internal/msan
[TOOLEXEC] Running compile for package internal/msan
```

### 環境を綺麗にする
このままだとビルドキャッシュが残っているため、別プログラムをビルドする際にキャッシュヒットすると他のプログラムでもプリントされてしまいます。

一度下記のコマンドでキャッシュを削除することをお勧めします

```
go clean -cache
```

### この観察から分かること

実行結果を見ると：
1. **compile が先、link が後**: .go → .o → 実行ファイルの流れ
2. **パッケージごとにコンパイル**: main パッケージは個別にコンパイルされる
3. **link は最後に1回だけ**: すべての .o ファイルをまとめる

---

## Step 3: ビルドにGopherを仕込む

### Go Gopher ASCII アートを作ろう

まずは、表示したい Gopher を準備しましょう。
思いつかない方は下記のGopherを使用してください。

```go
// メインの Gopher
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
```

### 実行して確認

#### 1. Gopher とプログレスバーを見ながらビルド
```bash
cd skeleton/step3
go build -o mytoolexec main.go
go build -toolexec="$PWD/mytoolexec" ../../testdata/sample.go -o sample
```

出力例：
```
=== Go Build with Gopher ===
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

コンパイル中 ██████████████████████████████ ✅
リンク中 ██████████████████████████████ ✅

🎉 ビルド完了！
```

---

## まとめ

## 実践的な応用

学んだ技術は以下のような場面で活用できます：

1. **ビルド時間の測定**
   - 各パッケージのコンパイル時間を記録
   - ボトルネックの特定

2. **セキュリティ強化**
   - デバッグ情報の自動除去
   - ビルド時のセキュリティチェック

3. **CI/CD パイプライン**
   - ビルド情報の自動埋め込み
   - 環境別の最適化設定

## 参考資料

- [Go build documentation](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [Go toolchain internals](https://go.dev/doc/toolchain)

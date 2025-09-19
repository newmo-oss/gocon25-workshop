author: sivchari
summary: Introduction to Go Assembly
id: assembly
categories: codelab,markdown
environments: Web
status: Published

# Go アセンブリ Codelab

このコードラボでは、Go のアセンブリを読み書きする方法を学習します。Go コードがどのように機械語に変換されるかを理解し、簡単な関数をアセンブリで実装できるようになります。

## 学習目標

- Go コードがどのようなアセンブリに変換されるか理解する
- Go アセンブリ記法の基本を習得する
- 簡単な関数をアセンブリで実装できるようになる

## 📋 進め方

1. 各ステップの「学習内容」と「ゴール」を読みます
2. Step 1 では `make` コマンドで生成されるアセンブリを観察します
3. Step 2 では `skeleton/step2/asm_amd64.s` の TODO を実装します
4. テストで動作を確認：`go test`
5. 詰まったら `solution/` の完成版を参照します

---

## Go のアセンブリについて

### Go アセンブリとは

Go は独自のアセンブリ記法を使用します：

- Plan 9 アセンブラをベースに Go 用に拡張された記法
- Intel や AT&T 記法とは異なる構文
- 疑似レジスタ（FP, SP, SB, PC）による抽象化
- プラットフォームに依存しない記述

### なぜ Go アセンブリを学ぶのか

1. **標準ライブラリの理解**: runtime, syncなどでは頻出します。特にgoroutineをはじめとしたGoランタイムの実装ではクロスプラットフォーム実現のためにほぼアセンブリで記述されています。
2. **Go の内部動作の理解**: コンパイラがどのようなコードを生成するか学ぶ

---

## Step 1: アセンブリを読む

### ゴール
Go コードから生成されるアセンブリを読んで理解できるようになる

### 学習内容

Go コンパイラが生成するアセンブリを観察し、以下を理解します：
- 関数の呼び出し規約
- レジスタの使い方
- スタックフレームの構造

### 基本的な命令

| 命令 | 説明 | 例 |
|------|------|-----|
| MOVQ | 64ビット値の移動 | `MOVQ AX, BX` (AX → BX) |
| ADDQ | 64ビット加算 | `ADDQ BX, AX` (AX += BX) |
| SUBQ | 64ビット減算 | `SUBQ BX, AX` (AX -= BX) |
| RET | 関数から戻る | `RET` |

### 疑似レジスタ

| レジスタ | 説明 |
|----------|------|
| FP | Frame Pointer - 引数と戻り値にアクセス |
| SP | Stack Pointer - ローカル変数にアクセス |
| SB | Static Base - グローバル変数にアクセス |
| PC | Program Counter - 次の命令のアドレス |

### 実践

`skeleton/step1/` ディレクトリで以下のコマンドを実行：

```bash
cd skeleton/step1

# add関数のアセンブリを見る
make add
```

### 観察ポイント

#### 1. add関数
```go
func add(a, b int) int {
    return a + b
}
```

実際に生成されるアセンブリ（最適化なし `-N -l`）：
```asm
main.add STEXT nosplit size=39 args=0x10 locals=0x10
    TEXT    main.add(SB), NOSPLIT|ABIInternal, $16-16
    PUSHQ   BP                    // ベースポインタを保存
    MOVQ    SP, BP                // 現在のスタックポインタを保存
    SUBQ    $8, SP                // スタック領域を確保
    MOVQ    AX, main.a+24(SP)    // 引数a（AXレジスタ経由）をスタックに保存
    MOVQ    BX, main.b+32(SP)    // 引数b（BXレジスタ経由）をスタックに保存
    MOVQ    $0, main.~r0(SP)      // 戻り値領域を初期化
    ADDQ    BX, AX                // AX = AX + BX (実際の加算)
    MOVQ    AX, main.~r0(SP)      // 結果をスタックの戻り値領域に保存
    ADDQ    $8, SP                // スタック領域を解放
    POPQ    BP                    // ベースポインタを復元
    RET                           // 関数から戻る
```

**重要なポイント**：
- Go 1.17以降、ABIInternal により引数は AX, BX レジスタで渡される
- スタック操作（PUSHQ/POPQ）で関数の開始/終了を管理
- 最適化を無効にすると、中間的なスタック操作が見える

#### 2. sub関数
```go
func sub(a, b int) int {
    return a - b
}
```

生成されるアセンブリでは、ADDQ の代わりに SUBQ 命令が使われます：
```asm
main.sub STEXT nosplit size=39 args=0x10 locals=0x10
    TEXT    main.sub(SB), NOSPLIT|ABIInternal, $16-16
    // ... 前処理は add と同じ ...
    SUBQ    BX, AX                // AX = AX - BX (引き算)
    // ... 後処理は add と同じ ...
    RET
```

---

## Step 2: アセンブリで関数を書く

### ゴール
Go アセンブリで簡単な関数を実装できるようになる

### 学習内容

実際にアセンブリで関数を書いて、Go から呼び出す方法を学びます。

### 関数の構造

```asm
TEXT ·FuncName(SB), NOSPLIT, $0-24
    // 関数の実装
    RET
```

- `TEXT`: 関数定義の開始
- `·FuncName`: 関数名（中点 · に注意）
- `(SB)`: Static Base からの相対
- `NOSPLIT`: スタック拡張チェックをスキップ
- `$0-24`: スタックサイズ-引数と戻り値のサイズ

### 実装タスク

`skeleton/step2/asm_amd64.s` の TODO を実装してください。

#### 1. Add関数の実装
2つの int64 を足し算する関数を実装：

```asm
// func Add(a, b int64) int64
TEXT ·Add(SB), NOSPLIT, $0-24
    // TODO: 実装
    RET
```

ヒント：
- `a` は `a+0(FP)` でアクセス
- `b` は `b+8(FP)` でアクセス（int64 は 8バイト）
- 戻り値は `ret+16(FP)` に書き込む

#### 2. Sub関数の実装
2つの int64 を引き算する関数を実装：

```asm
// func Sub(a, b int64) int64
TEXT ·Sub(SB), NOSPLIT, $0-24
    // TODO: 実装
    RET
```

ヒント：
- `SUBQ` 命令で引き算（AX = AX - BX）


### テストの実行

```bash
cd skeleton/step2

# 各関数のテスト
go test -v -run TestAdd
go test -v -run TestSub

# すべてのテスト
go test -v
```

---

## まとめ

このコードラボで学んだこと：

### 📚 Step 1: アセンブリを読む
- Go コンパイラが生成するアセンブリの構造
- Go アセンブリ記法の基本（MOVQ, ADDQ, RET など）
- 疑似レジスタ（FP, SP）の役割

### 📚 Step 2: アセンブリを書く
- TEXT ディレクティブによる関数定義
- 引数と戻り値へのアクセス方法
- 基本的な算術演算（加算・減算）の実装

## 実践的な応用

学んだ知識は以下の場面で活用できます：

- **標準ライブラリの理解**
   - Go schedulerの実装理解
   - math、syncのようなランタイムレベルでの操作を必要とするライブラリの理解


## 参考資料

- [A Quick Guide to Go's Assembler](https://go.dev/doc/asm) - Go 公式のアセンブリガイド

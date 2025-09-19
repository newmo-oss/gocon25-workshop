package main

// 1. シンプルな足し算
// この関数がどのようなアセンブリになるか観察します
func add(a, b int) int {
	return a + b
}

// 2. シンプルな引き算
// 引き算がどのようにアセンブリに変換されるか観察します
func sub(a, b int) int {
	return a - b
}

func main() {
	_ = add(3, 4)
	_ = sub(10, 3)
}

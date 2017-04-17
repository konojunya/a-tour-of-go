package main

import "fmt"

// defer ステートメントは、deferへ渡した関数の実行を呼び出し元の関数の終わり(returnする)まで遅延させるものです。
// defer へ渡した関数の引数は、すぐに評価されますが、その関数自体は呼び出し元の関数がreturnするまで実行されません

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

package main

import "fmt"

func main() {
  i ,j := 42,2701

  p := &i // iのポインタをpへ渡す
  fmt.Println(*p) // ポインタを通してiを呼ぶ
  *p = 21 // ポインタを通してiに値を代入する
  fmt.Println(i) // iの値を確認する

  p = &j // jのポインタをpへ渡す
  *p = *p / 37 // ポインタを通してjを割る
  fmt.Println(j)
}
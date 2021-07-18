package main

import (
	"fmt"
)

func main() {
	NUM := 1
	Num := 2

	fmt.Println(NUM)
	fmt.Println(Num)
}

// ●1文字目が小文字の場合はそのパッケージだけで使える変数や関数
// ●1文字目が大文字の場合は他のパッケージから使える変数や関数になる。
// ●Printlnは1文字目が大文字になっており、fmtのパッケージ関数になっていてmainパッケージの関数ではないが使える理由は1文字目が大文字になっているから。
// ●class, return, while, forなどは変数として使用できない。
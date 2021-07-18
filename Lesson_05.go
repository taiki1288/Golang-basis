package main

import (
	"fmt"
)

func main() {
	num01 := 123
	var num02 int = 1234567890
	// ↑int型を宣言
	num03 := 1.23
	// ↑データ型を省略
	var num04 float64 = 1.23456789
	// ↑float64型を宣言

	fmt.Println(num01)
	fmt.Println(num02)
	fmt.Println(num03)
	fmt.Println(num04)
}

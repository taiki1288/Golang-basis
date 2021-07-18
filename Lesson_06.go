package main

import (
  "fmt"
  "reflect"
  //↑reflect型の宣言
)
  
  
func main() {
	num01 := 123
	var num02 int = 1234567890
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(reflect.TypeOf(num01))
	fmt.Println(reflect.TypeOf(num02))
	fmt.Println(reflect.TypeOf(num03))
	fmt.Println(reflect.TypeOf(num04))
}

// ●Goのデータ型はreflectのパッケージのtypeofを使用して確認することができる。

package main

import (
  "fmt"
  "reflect"
  //↑reflect型の宣言
)

func main() {
	a := 10
	b := 1
	var num_bool bool = a > b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))
}

// ●↑bool型
// ●fmt.Println(num_bool)の結果でtrueと表示される
// ●fmt.Println(reflect.TypeOf(num_bool))の結果でbool型のboolと表示される。
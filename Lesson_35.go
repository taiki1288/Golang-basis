package main 

import "fmt"

type Student struct {
	name string
	math,english float64
}

func main () {
    s := Student{"sato", 80, 70}
	fmt.Println(s)
}

//●フィールドの順番通りに値を代入しなければエラーが起こる。
//●一部のフィールドに値を代入することができない。

// func main () {
// 	s := Student{name: "sato", math: 80, english: 70}
// 	fmt.Println(s)
// }
// ↑このようにフィールド名を指定することでエラーは起きない。
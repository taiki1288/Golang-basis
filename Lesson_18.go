package main

import "fmt"

func main() {

	x := 10
	y := 12

	if age := x + y ; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}

// ●簡易文はifの中だけでしか使えない。
package main

import "fmt"

func main () {
	x := 8
	y := 12
	z := 20

	x += 10
	z += y

	fmt.Println(x)
	fmt.Println(z)
}
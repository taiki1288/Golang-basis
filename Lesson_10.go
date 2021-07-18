package main

import "fmt"

func main() {
	a := [...]string{"yamashita", "yamamoto", "suzuki"}
	a[1] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}

// ●[要素数]をドットで省略することもできる。
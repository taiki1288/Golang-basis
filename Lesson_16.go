package main

import "fmt"

func main() {
	age := 0

	if age >= 20 {
		fmt.Println("adult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
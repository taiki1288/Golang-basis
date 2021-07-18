package main

import "fmt"

func main () {
	x := 7

	if x >= 5 && x < 10 {
		fmt.Println("５以上10未満")
	} else {
		fmt.Println("それ以外")
	}

}
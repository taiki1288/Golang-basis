package main

import "fmt"

func main () {
	i := 0
	for {
		fmt.Println(i)
		if i == 4 {
			break
		}
		i++
	}
}
package main 

import "fmt"

func main () {
	arr := [...]int{2, 4, 6, 8, 10}

	for i := 0; i<= 4; i++ {
		fmt.Println(arr[i])
	}
}
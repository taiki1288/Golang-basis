package main

import "fmt"

type User struct {
	name string
}

func (b User) cal(weight, height float64) (result float64) {
	result = weight / height /height * 10000
	return
}

func main () {
	user01 := User{name: "taiki"}
	fmt.Println(user01.name, user01.cal(65, 172))
}

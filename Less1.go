package main

import "fmt"

type employee string

type Human struct{
	name string
	age int
}

func main(){
	human := Human{
		name: "John",
		age: 17,
	}
	fmt.Println(human)
}

package mylib

import "fmt"

type person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
	person := person{Name: "Mike", Age: 20}
	fmt.Println(person)
}

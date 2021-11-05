package main

import "fmt"

type Myint int

func (i Myint) double() int {
	return int(i) * 2 //castしないとmain.Myint型を返すことになる
}

func main() {
	myInt := Myint(10)
	fmt.Println(myInt)
	fmt.Println(myInt.double())
}

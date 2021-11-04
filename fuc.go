package main

import "fmt"

func add1(x int, y int) {
	fmt.Println("add1 function")
	fmt.Println(x + y)
}

func add2(x int, y int) int {
	fmt.Println("add2 function")
	return x + y
}

func add3(x int, y int) (int, int) {
	fmt.Println("add2 function")
	return x + y, x - y
}

func calc(price, num int) (result int) {
	result = price * num
	return //result
}

func main() {
	add1(1, 2)
	r := add2(10, 20)
	fmt.Println(r)
	r2, r3 := add3(100, 200)
	fmt.Println(r2, r3)
	r4 := calc(100, 3)
	fmt.Println(r4)

	f := func(x int) {
		fmt.Println("Inner func", x)
	}
	f(100)

	func(x int) {
		fmt.Println("Inner func", x)
	}(300)
}

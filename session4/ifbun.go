package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	}
	return "no"
}

func main() {
	/*
		num := 6
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	/*
		result := by2(10)
		if result == "ok" {
			fmt.Println("great")
		}
		fmt.Println(result)
	*/

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	//fmt.Println(result2)

	x, y := 10, 20
	if x == 10 && y == 10 {
		fmt.Println("&&")
	} else if x == 10 || y == 10 {
		fmt.Println("||")
	}
}

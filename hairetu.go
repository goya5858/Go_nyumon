package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	n := []int{1, 2, 3, 4, 5, 6, 7}
	n = append(n, 10)
	n[2] = 100
	fmt.Println(n)

	var board [][]int = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)
}

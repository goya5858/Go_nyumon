package main

import "fmt"

func main() {
	var p *int = new(int) //メモリを確保
	fmt.Println(p)

	var p2 *int // メモリは確保してない
	fmt.Println(p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)
}

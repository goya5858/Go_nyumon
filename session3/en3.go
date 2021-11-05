package main

import "fmt"

func m() {
	m2 := map[string]int{"Mile": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T, %v", m2, m2)
}

func main() {
	m()
}

package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["Nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["apple"] = 100
	fmt.Println(m2)

	//var m3 map[string]int
	//m3["apple"] = 100

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}

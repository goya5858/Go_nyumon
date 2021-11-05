package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	//v.X  =1000 //これでも正常に動作する
	(*v).X = 1000
}

func main() {
	v := Vertex{100, 1000, "test"}
	changeVertex(v)
	fmt.Println(v)

	changeVertex2(&v)
	fmt.Println(v)
	/*
		v := Vertex{X: 1, Y: 2}
		fmt.Println(v)
		v.X = 100
		fmt.Println(v.X, v.Y)

		v2 := Vertex{X: 1, S: "test"}
		fmt.Println(v2)

		var v3 Vertex
		fmt.Println(v3)

		v4 := new(Vertex)
		fmt.Println(v4)

		v5 := &Vertex{}
		fmt.Println(v5)
	*/
}

package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}
func (v *Vertex) Scale(i int) {
	v.x *= i
	v.y *= i
}

//実質的なコンストラクタ
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x *= i
	v.y *= i
	v.z *= i
}

//実質的なコンストラクタ
func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

/*
func Area(v Vertex) int {
	return v.X * v.Y
}
*/

func main() {
	//v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v := New3D(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}

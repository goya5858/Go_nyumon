package main

import "fmt"

// 返り値に「関数の定義」を定義する
// 今回の場合は "返り値にintを持つ関数" を返り値に定義してる
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// 返り値は func(radius float64) float64 の関数
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return radius * radius * pi
	}
}

func main() {
	// ここで実体化されてるcounterは関数のこと
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	//pi=3.14で設定
	f1 := circleArea(3.14)
	fmt.Println(f1(10))
	//pi=3.0で設定
	f2 := circleArea(3.0)
	fmt.Println(f2(10))
}

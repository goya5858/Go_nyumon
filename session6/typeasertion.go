package main

import "fmt"

func do(i interface{}) {
	/*
		ii := i.(int) //タイプアサーション
		// interface -> 別のタイプに変換　をタイプアサーションという
		ii *= 2
		fmt.Println(ii)
	*/
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I dont know : %T\n", v)
	}
}

func main() {
	//do(10)
	// ↑↓は同じこと　空のinterfaceにintを入れてる
	//var i interface{} = 30
	//do(i)
	do(10)
	do("mike")
	do(true)
}

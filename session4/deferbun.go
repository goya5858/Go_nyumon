package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("Hello foo")
}

func main() {
	/*
		defer fmt.Println("world") //関数内の他の処理が終わった後に実行
		foo()
		fmt.Println("Hwllo ")
	*/
	fmt.Println("run")
	defer fmt.Println(1) //上ほど最後に実行される
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	/*
		file, _ := os.Open("./lesson.go")
		defer file.Close()
		data := make([]byte, 100)
		file.Read(data)
		fmt.Println(string(data))
	*/
}

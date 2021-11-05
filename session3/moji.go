package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello" + "World")
	fmt.Println(string("Hello World"[0])) //キャスト
	s := "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
}

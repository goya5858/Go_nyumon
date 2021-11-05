package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./en1.go")

	//エラーハンドリングのコード
	if err != nil {
		log.Fatal("Error")
	}
	defer file.Close()

	data := make([]byte, 100)

	// errが上書きされる　:=　は新規で定義できるものが１つでもあれば使える
	count, err := file.Read(data)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println(count, string(data))
}

package main

import (
	"fmt"
	"sync"
)

func Normal(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //wgに並列処理の作業が１つあることを教える
	go goroutine("world", &wg)
	Normal("hello")
	//time.Sleep(2000 * time.Millisecond)
	wg.Wait() //登録されてる回数分Doneがされるのを待つ
}

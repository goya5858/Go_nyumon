package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process : ", i)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		go producer(ch, i)
		wg.Add(1)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
}

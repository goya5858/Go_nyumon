package main

import (
	"fmt"
	"time"
)

func main() {
	// TickもAfterも、共にchanを返す関数
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OUTROOP:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			break OUTROOP
			//return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("################")
}

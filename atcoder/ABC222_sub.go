package main

import (
	"fmt"
	"strings"
)

func main() {
	var N string
	fmt.Scan(&N)

	delta := 4 - len(N)
	if delta >= 0 {
		for i := 0; i < delta; i++ {
			fmt.Printf("%v", 0)
		}
	}
	s_N := strings.Split(N, "")
	for _, s := range s_N {
		fmt.Printf("%v", s)
	}
	fmt.Printf("\n")
}

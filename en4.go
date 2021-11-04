package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var num_max int = 0
	for _, item := range l {
		if num_max < item {
			num_max = item
		}
	}
	fmt.Println(num_max)
}

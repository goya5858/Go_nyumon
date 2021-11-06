package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	inputs := make([]int, n) //n個分の配列を作成
	for i := 0; i < n; i++ {
		fmt.Scan(&inputs[i])
	}
	fmt.Println(inputs)
}

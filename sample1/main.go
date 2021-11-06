package main

import (
	"fmt"
	print "sample1/formatter"
	"sample1/math"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

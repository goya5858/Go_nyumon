package main

import (
	"fmt"
	"time"
)

func getosName() string {
	return "mac"
}

func main() {
	//os := getosName()
	switch os := getosName(); os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("default")
	}
	//fmt.Println(os)

	t := time.Now()
	fmt.Println(t.Hour())

	switch { //条件を書かなくてもいい
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("AfterNoon")
	}
}

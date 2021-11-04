package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("hellow", time.Now())
	fmt.Println(user.Current())
}

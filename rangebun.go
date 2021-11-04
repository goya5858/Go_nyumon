package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banan": 200, "melon": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

}

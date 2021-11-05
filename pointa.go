package main

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	println(n)
	/*
		fmt.Println(n)
		fmt.Println(&n)

		var p *int = &n
		fmt.Println(p)
		fmt.Println(*p)
	*/
}

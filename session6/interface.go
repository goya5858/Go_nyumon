package main

import "fmt"

//Human interfaceを設定
//Humanから作られるオブジェクトはSayメソッドを持つ必要がある
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run!")
	} else {
		fmt.Println("getout")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	//fmt.Println(mike.Say())
	DriveCar(mike)
}

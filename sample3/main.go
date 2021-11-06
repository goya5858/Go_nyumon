package main

import (
	"fmt"
	"sample3/person"
)

func main() {
	bob := person.Person{PetName: "Flutty"}
	fmt.Println(bob.Pet())
}

/*
以下のエラーを想定しています
循環参照です

package command-line-arguments
        imports sample3/person
        imports sample3/pet
        imports sample3/person: import cycle not allowed
*/

package main

import (
	"encoding/json"
	"example/hello/person"
	"fmt"
)

func main() {
	pessoa := person.Person{
		Name: "Linus",
		Lastname: "Torvalds",
		Age: 54,
	}

	personString, _ := json.Marshal(person.CreatePerson(pessoa))
	
	fmt.Println(string(personString))
}
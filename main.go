package main

import (
	"encoding/json"
	"example/hello/person"
	"fmt"
)

func main() {
	pessoa := person.Person{
		Name: "Michele",
		Lastname: "Martins",
		Age: 33,
	}

	personString, _ := json.Marshal(person.CreatePerson(pessoa))
	
	fmt.Println(string(personString))
}
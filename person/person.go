package person

type Person struct {
	Name     string
	Age      int
	Lastname string
}

func CreatePerson(person Person) Person {
	return person
}
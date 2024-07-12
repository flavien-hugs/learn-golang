package main

import "fmt"

func main() {
	firstPerson := person{}
	firstPerson.name = "Flavien HUGS"
	firstPerson.size = 1.75

	// display value
	fmt.Println("name:", firstPerson.name, "\nsize:", firstPerson.size)

	// OR
	secondPerson := person{"Ulrich GERO", 1.80}
	fmt.Println("name:", secondPerson.name, "\nsize:", secondPerson.size)
}

type person struct {
	name string
	size float32
}

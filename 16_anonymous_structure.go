package main

import (
	"fmt"
	"log"
)

func main() {

	person := struct {
		name string
		age int
		degree string
	}{
		name: "Josué Parra",
		age: 4,
		degree: "Student",
	}

	addr := address { "Armando Carrera", 5218}

	fmt.Println(person)

	fmt.Println(addr)

	log.Printf("Person - name: %s - age: %d - degree: %s", person.name, person.age, person.degree)
}


// Anonymous Fields

type address struct {
	string
	int
}
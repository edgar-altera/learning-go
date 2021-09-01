package main

import "fmt"

func main() {

	var p1 person

	p1.name = "Edgar Parra"
	
	p1.name = "Developer"

	p1.address = address { "Santiago", "Armando Carrera", 5218 }

	fmt.Println(p1.address.street)
}

type person struct {
	name string
	degree string
	address address
}

type address struct {
	city string
	street string
	number int
}
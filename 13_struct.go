package main

import "fmt"

func main() {

	var p1 person
	
	p1.name = "Josué"

	p2 := person { "Edgar", 35, "Developer", true }

	p3 := person { name: "Rebeca", age: 1}

	fmt.Println(p1.name)
	fmt.Println(p2.name)
	fmt.Println(p3.name)

	fmt.Printf("%+v", p3)
}

type person struct {
	name string
	age int
	degree string
	status bool
}

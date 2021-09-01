package main

import "fmt"

func main() {

	p1 := Person { "Edgar", 35 }

	fmt.Println(p1)

	setValues(&p1)

	fmt.Println(p1)
}

type Person struct {
	name string
	age int
}

func setValues(p *Person) {

	p.name = "Josué"
	p.age = 4
}
package main

import "fmt"

func main() {

	number := 100

	p1 := &number

	p2 := &p1

	fmt.Printf("Number: %d \n", number)
	fmt.Printf("Pointer 1: %X \n", p1)
	fmt.Printf("Pointer 2: %X \n", p2)

	fmt.Printf("Valor de number a través de p1: %d \n", *p1)
	fmt.Printf("Valor de number a través de p2: %d \n", **p2)

	**p2 = 200

	fmt.Printf("Number: %d \n", number)

	// Comparing pointers

	x := &number
	y := &number
	var z *int
	var zz *int

	fmt.Println(x == y)
	fmt.Println(x == z)
	fmt.Println(x != z)
	fmt.Println(z == zz)
}

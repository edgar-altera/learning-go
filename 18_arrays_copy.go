package main

import "fmt"

func main() {

	// Valor

	a1 := [...] int { 1, 2, 3 }
	a2 := a1

	a2[0] = 9

	fmt.Println("Array 1: ", a1)
	fmt.Println("Array 2: ", a2)

	// Referencia
	
	a3 := &a1

	fmt.Println("Array 3: ", a3)

	a1[0] = 35

	fmt.Println("Array 1: ", a1)
	fmt.Println("Array 3: ", a3)
}
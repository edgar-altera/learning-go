package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "Hello \n World!"

	// raw literals
	s2 := `Hola \n Mundo!`

	var multi_lines string

	multi_lines = `
		Primera linea
		Segunda linea
		.
		.
		.
	`

	var s3 string = "Hi!"

	s3 = "Hola"

	// Strings are inmutable
	// s3[0] = 'i'; --> Error --> cannot assign to s3[0] (strings are immutable)

	fmt.Printf("Cadena 1: %s \n", s1)

	fmt.Printf("Cadena 2: %s \n", s2)

	fmt.Printf("Cadena 3: %s \n", s3)

	fmt.Printf("Cadena multi linea: %s \n", multi_lines)


	city := "Santiago de Chile"

	for k, v := range city {

		fmt.Printf("Key: %d - Value: %c \n", k, v)
	}

	for i := 0; i < len(city); i++ {

		fmt.Printf("Character: %c - utf8 value: %v \n", city[i], city[i])
	}

	fmt.Printf("Len function: %d \n", len(city))

	fmt.Printf("RuneCountInString  function: %d \n", utf8.RuneCountInString(city))
}
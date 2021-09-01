package main

import "fmt"

func main() {

	// Array string
	var langs [3] string

	langs[0] = "PHP"
	langs[1] = "JavaScript"
	langs[2] = "Go"

	for key, lang := range langs {

		fmt.Printf("Array Key: %d - Array Value: %s \n", key, lang)
	}

	// Array ints

	numbers := [...] int {1, 2, 3, 4, 5, 6, 7, 8, 9 }

	total := 0

	for i := 0; i < len(numbers); i++ {

		total += i
	}

	fmt.Printf("La suma de los numeros del array es: %d \n", total)

	// Print X 

	n := 20

	for i := 0; i < n; i++ {
		
		for j := 0; j < n; j++ {

			if ( (i == j) || ((i + j) == (n - 1))) {
				
				fmt.Printf("X")

			} else {

				fmt.Printf("_")
			}

			if ( j == ( n - 1)) {

				fmt.Printf("\n")
			}
		}
	}

	// Array Multi Dimensional
	
	multi := [3][3] string {
		{ "HTML", "CSS", "JavaScript"},
		{ "PHP", "Python", "Go"},
		{ "Laravel", "Vue", "React"},
	}

	fmt.Println(multi)

	// ... Dimension definida por la inicialización

	letters := [...] string { "A", "B", "C"}

	fmt.Printf("El array %v es de longitud %d \n", letters, len(letters))

	// Comparar matrices

	aa := [3] int { 1, 2, 3 }
	ab := [...] int { 1, 2, 3 }
	ac := [...] int { 1, 2, 33 }

	fmt.Printf("aa == ab --> %t \n", aa == ab)
	fmt.Printf("aa == ac --> %t \n", aa == ac)
}
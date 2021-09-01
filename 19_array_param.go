package main

import "fmt"

func main() {

	n := [...] int { 20, 18, 19, 16 }

	fmt.Printf("El promedio de notas es: %d \n", avg(n))
}

func avg(notes [4] int) int {

	total := 0

	size := len(notes)

	for i := 0; i < size; i++ {
		
		total += notes[i]
	}

	average := total / size

	return average
}
package main

import "fmt"

func main() {

	nombresArray := [4] string {"Ana", "José", "Daniel", "María"}
	nombresSlice := nombresArray[0:3]

	for key, value := range nombresSlice {

		fmt.Printf("Key: %d - Value: %s \n", key, value)
	}

	fmt.Printf("Array Len: %d Cap: %d \n", len(nombresSlice), cap(nombresSlice))

	nombresSlice = append(nombresSlice, "Pedro")
	nombresSlice = append(nombresSlice, "Pablo") 

	for key, value := range nombresSlice {

		fmt.Printf("Key: %d - Value: %s \n", key, value)
	}

	fmt.Printf("Array Len: %d Cap: %d \n", len(nombresSlice), cap(nombresSlice))


	array := [3] int { 1, 2, 3}
	slice := array[:2]

	fmt.Println("Original Array: ", array)
	fmt.Println("Original Slice: ", slice)

	array[0] = 9
	slice[1] = 8

	fmt.Println("New Array: ", array)
	fmt.Println("New Slice: ", slice)
}
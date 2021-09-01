package main

import "fmt"

func main() {

	sum(2, 2)

	defer sum(5, 5)

	sum(10, 40)
	
	defer sum(100, 200)

	defer sum(15, 15);
}

func sum(a, b int) {

	result := a + b

	fmt.Printf("%d + %d = %d \n", a, b, result)
}
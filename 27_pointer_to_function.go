package main

import "fmt"

func main() {

	number := 375

	fmt.Printf("Number value: %d \n", number)

	ptf(&number)

	fmt.Printf("Number value: %d \n", number)
}

func ptf(p *int) {

	*p = 500
}
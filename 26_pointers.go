package main

import "fmt"

func main() {

	x := 0xFF
    y := 0x9C
     
    // Displaying the values
    fmt.Printf("Type of variable x is %T\n", x)
    fmt.Printf("Value of x in hexadecimal is %X\n", x)
    fmt.Printf("Value of x in decimal is %v\n", x)
     
    fmt.Printf("Type of variable y is %T\n", y)
    fmt.Printf("Value of y in hexadecimal is %X\n", y)
    fmt.Printf("Value of y in decimal is %v\n", y)   


	a := 1900
	var b *int
	var c *int

	b = &a

	fmt.Printf("Valor de a: %d \n", a)
	fmt.Printf("Valor de b: %v \n", b)
	fmt.Printf("Valor de a, através de b: %d \n", *b) // desreferenciación 
	fmt.Printf("Valor de c: %v \n", c)

	*b = 1950

	fmt.Printf("Valor de a: %d \n", a)
	fmt.Printf("Valor de a atraves de b: %d \n", *b)
}
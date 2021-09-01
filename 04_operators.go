package main

import "fmt"

func main() {

	x := 5
	y := 4

	fmt.Printf("%d != %d: %t \n", x, y, x != y)
	fmt.Printf("%d == %d: %t \n", x, y, x == y)
	fmt.Printf("%d > %d: %t \n", x, y, x > y)
	fmt.Printf("%d < %d: %t \n", x, y, x < y)
	fmt.Printf("%d >= %d: %t \n", x, y, x >= y)
	fmt.Printf("%d <= %d: %t \n", x, y, x <= y)
}
package main

import "fmt"

func main() {

	a, b := 5, 20

	m, _ := mul_div(a, b)

	fmt.Printf("%d * %d = %d \n", a, b, m)
}

func mul_div(x, y int) (int, int) {

	mul := x * y

	div := x / y

	return mul, div
}

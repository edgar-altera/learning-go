package main

import "fmt"

func main() {

	result := area(12, 10)

	fmt.Printf("El area del rectangulo es: %d \n", result)

	a, b := mul_div(20, 2)

	fmt.Printf("20 * 2 = %d | 20 / 2 = %d \n", a, b)

	p := 100

	q := 300

	fmt.Printf("p = %d and q=%d \n", p, q)

	swap(&p, &q)

	fmt.Printf("p = %d and q=%d \n", p, q)
}

func area(width, length int) int {

	value := width * length

	return value
}

func mul_div(x, y int) (int, int) {
	
	mul := x * y
	
	div := x / y
	
	return mul, div
}

func swap(a, b *int) {

	tmp:= *a

	*a = *b

	*b = tmp
}


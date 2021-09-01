package main

import (
	"fmt"
	"strings"
)

func main() {
	
	fmt.Println(joinstr())

	fmt.Println(joinstr("Hello"))

	fmt.Println(joinstr("Coca", "Cola"))
	
	fmt.Println(joinstr("A", "B", "C"))
	
	fmt.Println(joinstr("1", "2", "3", "4", "5"))

	numbers := [] string { "100", "200", "300", "400", "500", "600", "700" }

	fmt.Println(joinstr(numbers...))

	fmt.Printf("Suma = %d \n", sum(100, 500, 2000, 34, 350))
}

func joinstr(element... string) string {

	return strings.Join(element, "-")

}

func sum(numbers... int) int {

	total := 0

	for i := 0; i < len(numbers); i++ {
		
		total += numbers[i]
	}

	return total
}
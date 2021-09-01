package main

import (
	"fmt"
	"sort"
)

func main() {

	numbers := [] int { 1, 3, 2 }

	isOrder := sort.IntsAreSorted(numbers)

	fmt.Printf("Original Array: %v \n", numbers)
	fmt.Printf("Array sorted = %t \n", isOrder)

	sort.Ints(numbers)

	isOrder = sort.IntsAreSorted(numbers)

	fmt.Printf("Sorted Array: %v \n", numbers)
	fmt.Printf("Array sorted = %t \n", isOrder)

	names := [] string { "Pedro", "Ana", "Carlos" }

	isOrder = sort.StringsAreSorted(names)

	fmt.Printf("Original Array: %v \n", names)
	fmt.Printf("Array sorted = %t \n", isOrder)

	sort.Strings(names)

	isOrder = sort.StringsAreSorted(names)
	
	fmt.Printf("Sorted Array: %v \n", names)
	fmt.Printf("Array sorted = %t \n", isOrder)
}
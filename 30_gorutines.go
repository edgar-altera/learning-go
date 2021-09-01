package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hi!")

	for i := 0; i < 5; i++ {

		go func(n int) {
			fmt.Printf("I = %d \n", n)
		}(i)
	}

	time.Sleep(1e9)

	fmt.Println("Bye!")
}


package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome!")

	go func() {
		fmt.Println("Hi!")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("Goodbye")
}
package main

import "fmt"

func main() {

	var channel1 chan int

	fmt.Println("Valor de channel1: ", channel1)
	fmt.Printf("Tipo de channel1: %T \n", channel1)	

	channel2 := make(chan int)

	fmt.Println("Valor de channel2: ", channel2)
	fmt.Printf("Tipo de channel2: %T \n", channel2)	

	go add(channel2)

	channel2 <- 50 // Send

	fmt.Println("Bye")
}

func add(ch chan int) {

	number := <-ch // Receive
	
	fmt.Println(250 + number)
}
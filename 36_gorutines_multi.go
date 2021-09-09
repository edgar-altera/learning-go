package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start")

	go ids()

	go names()

	time.Sleep(10 * time.Second)

}

func ids() {

	data := [] int { 10, 20, 30, 40, 50 }

	for i := 0; i < 5; i++ {
		
		time.Sleep(1000 * time.Millisecond)

		fmt.Printf("i: %d id: %d \n", i, data[i])
	}
}

func names() {

	data := [] string { "Ana", "Pedro", "Pablo", "Maria", "Carlos" }

	for i := 0; i < 5; i++ {
		
		time.Sleep(1150 * time.Millisecond)

		fmt.Printf("i: %d name: %s \n", i, data[i])
	}
}
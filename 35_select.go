package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go channel1(c1)
	go channel2(c2)

	select {

		case msg := <- c1:
			fmt.Println(msg)
		
		case msg := <- c2:
			fmt.Println(msg)
	}

	time.Sleep(5 * time.Second)
}

func channel1(c1 chan string) {

	time.Sleep(2 * time.Second)

	c1 <- "Welcome to channel 1"
}

func channel2(c2 chan string) {

	time.Sleep(2 * time.Second)

	c2 <- "Welcome to channel 2"
}

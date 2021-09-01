package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := make(chan string)

	go display(c)

	for {

		resp, open := <- c

		if !open {

			fmt.Println("Channel close")
			break
		}

		fmt.Printf("Channel open. Resp: %s \n", resp)

	}
}

func display(channel chan string) {

	for i := 0; i < 5; i++ {

		channel <- "Hello World i: " + strconv.Itoa(i)
	}

	close(channel)
}

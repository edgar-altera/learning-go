package main

import "fmt"

func main() {

	ch := make(chan string)

	go addMessageToChannel(ch)

	for resp := range ch {

		fmt.Println(resp)
	}
}

func addMessageToChannel(channel chan string) {

	channel <- "Hola"
	channel <- "Hi"
	channel <- "Hello"

	close(channel)
}

package main

import "fmt"

func main() {

	channel_receive := make(<-chan string)

	channel_send := make(chan<- string)
	
	fmt.Printf("Receive channel: %T \n", channel_receive)

	fmt.Printf("Send channel: %T \n", channel_send)
}

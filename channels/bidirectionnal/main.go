package main

import (
	"fmt"
	"time"
)

// Receives a message and send it back
func cycle(c chan string) {
	var message string
	for {
		fmt.Println("Waiting for message")
		message = <-c
		fmt.Println(message)
		c <- message
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan string)
	go cycle(c)
	c <- "Hello World"
	fmt.Scanln()
	fmt.Println("Exit")
}

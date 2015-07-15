package main

import (
	"fmt"
	"time"
)

// Sets a function that takes a channel as argument. Only allows sending for
// this particular channel. The sending operation is blocking until someone
// reads the sent data.
func ping(c chan<- string) {
	for {
		c <- "ping"
	}
}

// Sets a function that takes a channel as argument. Only allows receiving for
// this particular channel. The receiving operation is blocking until some data
// is received.
func pong(c <-chan string) {
	for {
		<-c
		fmt.Println("pong")
		time.Sleep(time.Second * 1)
	}
}

// Create a string channel and starts two goroutines for the ping and pong
// functions. Waits for an input to stop the program.
func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	// Press enter to exit the program
	fmt.Scanln()
	print("Exit")
}

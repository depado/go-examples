package main

import (
	"flag"
	"fmt"
	"log"
)

var debug bool

// Prints a string using log.Println if and only if the debug flag was passed.
// t : The string to be printed.
func debugln(t string) {
	if debug {
		log.Println(t)
	}
}

func main() {
	flag.BoolVar(&debug, "debug", false, "Activates the debug logs")
	flag.Parse()

	debugln("This message will be displayed only if the -debug argument is used.")
	fmt.Println("This message will always be displayed.")
}

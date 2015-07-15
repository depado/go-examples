package main

import (
	"bufio"
	"fmt"
	"os"
)

// simpleReadLine scans only one time on os.Stdin. It then checks for errors
// and returns.
func simpleReadLine() (l string, err error) {
	rl := bufio.NewScanner(os.Stdin)
	rl.Scan()
	if err = rl.Err(); err != nil {
		return
	}
	l = rl.Text()
	return
}

func main() {
	var err error
	var line string

	fmt.Print("Enter something : ")
	line, err = simpleReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Println("You typed :", line)
}

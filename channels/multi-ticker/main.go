package main

import (
	"fmt"
	"log"
	"time"
)

type closableTicker struct {
	ticker *time.Ticker
	halt   chan bool
}

func (ct *closableTicker) stop() {
	ct.ticker.Stop()
	close(ct.halt)
}

// Takes one chan as argument (in this case a time.Ticker)
// A goroutine is started upon call which will listen on the given chan and send
// The data to the two other chans. This allows multiple chans to listen on the
// same source. Of course this whole mecanism will break if, later in the program,
// something tries to read from the source channel directly.
func multiTicker(ct closableTicker) (chan time.Time, chan time.Time) {
	a := make(chan time.Time)
	b := make(chan time.Time)
	go func() {
		for {
			select {
			case t := <-ct.ticker.C:
				a <- t
				b <- t
			case <-ct.halt:
				close(a)
				close(b)
				return
			}
		}
	}()
	return a, b
}

func doThing(c <-chan time.Time) {
	for range c {
		log.Println("Hello from doThing")
	}
	log.Println("Stopped doThing due to closed channel")
}

func doOtherThing(c <-chan time.Time) {
	for range c {
		log.Println("Hello from doOtherThing")
	}
	log.Println("Stopped doOtherThing due to closed channel")
}

func main() {
	ct := closableTicker{
		ticker: time.NewTicker(1 * time.Second),
		halt:   make(chan bool, 1),
	}
	a, b := multiTicker(ct)
	go doThing(a)
	go doOtherThing(b)
	time.Sleep(3 * time.Second)
	ct.stop()
	fmt.Scanln()
	log.Println("Exit")
}

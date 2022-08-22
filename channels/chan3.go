package main

import (
	"fmt"
)

func main() {
	c := mux(sendMessage("hi from c1: "), sendMessage("hi from c2: "))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func mux(in1, in2 <-chan string) <-chan string {
	// this takes two channels and spins up
	// a go routine for each one that reads from
	// the channel passed in and simply sends it out
	// on the channel created here - i.e. it multiplexes
	// the output of the channels passed as parameters
	// onto a single new channel - which this function
	// returns
	// the reason for this is to consume the messages
	// sent from the message generating goroutines as
	// they are available, rather than checking one
	// then checking the next (as in chan2.go)
	c := make(chan string)
	go func() {
		for {
			c <- <-in1 // read from channel in1, then send it to channel c
		}
	}()
	go func() {
		for {
			c <- <-in2 // read from channel in2, then send it to channel c
		}
	}()

	return c
}

func sendMessage(m string) <-chan string {
	// this function will create and return the channel on which
	// messages are sent, as well as start the goroutine that generates
	// the messages and puts them into the channel
	c := make(chan string)

	// below is a "function literal", the parens at the end are required for it to be called "in place"
	// NB no space between ending brace and parens
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", m, i)
		}
	}()

	return c
}

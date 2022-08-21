package main

import (
	"fmt"
)

func main() {
	c1 := sendMessage("hello!")
	c2 := sendMessage("hello!")

	for i := 0; i < 5; i++ {
		fmt.Printf("message c1 was: %s\n", <-c1)
		fmt.Printf("message c2 was: %s\n", <-c2)
	}
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

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go tx(c)
	go rx(c)
	time.Sleep(4 * time.Second)
}

func tx(c chan int) {
	c <- 1
	c <- 2
	c <- 3
}

func rx(c chan int) {
	fmt.Printf("val in chanel: %v\n", <-c)
	fmt.Printf("val in chanel: %v\n", <-c)
	fmt.Printf("val in chanel: %v\n", <-c)
}

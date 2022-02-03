package main

import (
	"fmt"
	"time"
)

// Can we find a way to keep inserting numbers even when the channel already has a value in it?

func channels() {
	c := make(chan int)
	done := make(chan struct{})

	go insertValues(c)
	go processValues(c, done)

	<-done
	fmt.Println("done!")
}

func insertValues(c chan<- int) {
	for i := 0; i < 1000; i++ {
		c <- i
		fmt.Printf("inserted %d into channel\n", i)
		time.Sleep(1 * time.Second)
	}

	close(c)
}

func processValues(c <-chan int, done chan<- struct{}) {
	for i := range c {
		fmt.Printf("received %d from channel\n", i)
		time.Sleep(3 * time.Second)
	}

	close(done)
}

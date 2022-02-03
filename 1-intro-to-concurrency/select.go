package main

import (
	"fmt"
	"time"
)

// Part 1: How can we exit the loop when both channels are done?
// Part 2: Can we have empty selects?

func usingSelect() {
	server1Channel := make(chan bool)
	server2Channel := make(chan int)

	go server1(server1Channel)
	go server2(server2Channel)

	for {
		select {
		case <-server1Channel:
			fmt.Println("server 1 receive")

		case n := <-server2Channel:
			fmt.Printf("server 2 received: %d\n", n)

		default:
			fmt.Println("nothing")
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("done!")
}

func server1(c chan<- bool) {
	for i := 0; i < 1000000; i++ {
		c <- true
		time.Sleep(1100 * time.Millisecond)
	}
}

func server2(c chan<- int) {
	for i := 0; i < 1000000; i++ {
		c <- i
		time.Sleep(600 * time.Millisecond)
	}
}

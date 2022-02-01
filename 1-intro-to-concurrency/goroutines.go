package main

import (
	"fmt"
	"time"
)

// Part 1: What if we put all of these in their own go-routines? (Fork-Join problem)
// Part 2: How can we wait for all the processes to finish?

// What about shared resources?

func goroutines() {
	now := time.Now()

	task1()
	task2()
	task3()
	task4()

	fmt.Println("time to complete all tasks:", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1 finished")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2 finished")
}

func task3() {
	fmt.Println("task 3 finished")
}

func task4() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("task 4 finished")
}

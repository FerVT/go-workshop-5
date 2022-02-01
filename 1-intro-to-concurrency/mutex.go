package main

import (
	"fmt"
	"sync"
)

// How can we guarantee that we will always get 1000 in the output?

func mutex() {
	counter := 0

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

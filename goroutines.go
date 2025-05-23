package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// keep track of all GRs
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Wait until this GR finishes before decrementing wg counter
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Index from GR 1:", i)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			fmt.Println("Index from GR 2:", j)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
	}()

	wg.Wait() // Wait for all GRs to finish before continuing execution
	fmt.Println("Done")

}

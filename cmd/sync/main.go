package main

import (
	"fmt"
	"sync"
)

func main() {
	// Each runner will execute 10 times
	printTimes := 10

	// Main WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(3)

	// Three WaitGroups to control the execution order between runners
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	wg2.Add(1) // Start blocked
	var wg3 sync.WaitGroup
	wg3.Add(1) // Start blocked

	// Runner 1: Starts immediately due to wg1.Wait() not blocking
	go func(id, printTimes int) {
		defer wg.Done()
		for i := 0; i < printTimes; i++ {
			wg1.Wait() // Wait for signal from Runner 3
			fmt.Printf("\nRunner %d \n", id)
			wg1.Add(1) // Protect against premature restart
			wg2.Done() // Signal Runner 2 to start
		}
	}(1, printTimes)

	// Runner 2: Starts after Runner 1 signals
	go func(id, printTimes int) {
		defer wg.Done()
		for i := 0; i < printTimes; i++ {
			wg2.Wait() // Wait for signal from Runner 1
			fmt.Printf("Runner %d \n", id)
			wg2.Add(1) // Protect against premature restart
			wg3.Done() // Signal Runner 3 to start
		}
	}(2, printTimes)

	// Runner 3: Starts after Runner 2 signals
	go func(id, printTimes int) {
		defer wg.Done()
		for i := 0; i < printTimes; i++ {
			wg3.Wait() // Wait for signal from Runner 2
			fmt.Printf("Runner %d \n", id)
			wg3.Add(1) // Protect against premature restart
			wg1.Done() // Signal Runner 1 to start
		}
	}(3, printTimes)

	// Wait for all goroutines to finish
	wg.Wait()
}

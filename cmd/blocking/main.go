package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	iter := 10
	chanRunner1 := make(chan struct{})
	chanRunner2 := make(chan struct{})

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= iter; i++ {
			<-chanRunner2
			chanRunner1 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i <= iter; i++ {
			<-chanRunner1
			chanRunner2 <- struct{}{}
		}
	}()

	wg.Wait()
	fmt.Println("finished")
}

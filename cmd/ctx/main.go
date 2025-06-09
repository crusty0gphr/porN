package main

import (
	"context"
	"fmt"
	"time"
)

func A(ctx context.Context) {
	// Create child context for B (cancels when A exits or manually canceled)
	ctxA, cancelA := context.WithCancel(ctx)
	defer cancelA() // Cancel B when A exits

	go B(ctxA) // Start B (no need to track separately)

	// A's work
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // Main cancelled us
			fmt.Println("A: shutdown signal received")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("A: working", i)

			// Example: A decides to cancel B after 3 iterations
			if i == 2 {
				fmt.Println("A: manually cancelling B")
				cancelA() // Cancel only B (not A itself)
			}
		}
	}
}

func B(ctx context.Context) {
	// Create child context for C
	ctxB, cancelB := context.WithCancel(ctx)
	defer cancelB() // Cancel C when B exits

	go C(ctxB) // Start C

	// B's work
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // A cancelled us
			fmt.Println("B: shutdown signal received")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("B: working", i)
		}
	}
}

func C(ctx context.Context) {
	// C's work
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // B cancelled us
			fmt.Println("C: shutdown signal received")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("C: working", i)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go A(ctx) // Start chain

	// Let run for 5 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("\nmain: cancelling entire chain")
	cancel() // Cancel everything

	// Allow time for shutdown messages
	time.Sleep(500 * time.Millisecond)
}

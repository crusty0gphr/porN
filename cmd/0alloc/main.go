package main

import (
	"fmt"
	"runtime"
)

// Stack is a fixed-size stack with a capacity of 16 elements.
type Stack struct {
	data [16]int
	top  int
}

// Push adds an element to the stack if there's space.
func (s *Stack) Push(x int) {
	if s.top < len(s.data)-1 {
		s.data[s.top+1] = x
		s.top++
	}
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() int {
	if s.top >= 0 {
		val := s.data[s.top]
		s.top--
		return val
	}
	return -1 // Error: empty stack
}

func main() {
	// Static memory allocation
	var s Stack
	s.top = -1

	// Record memory stats before operations
	var memStart, memEnd runtime.MemStats
	runtime.ReadMemStats(&memStart)

	// Perform operations without allocations
	for i := 0; i < 1000000; i++ {
		s.Push(i % 100)
	}
	for i := 0; i < 1000000; i++ {
		s.Pop()
	}

	// Record memory stats after operations
	runtime.ReadMemStats(&memEnd)

	// Output memory usage difference
	fmt.Printf("TotalAlloc delta: %v KiB\n", (memEnd.TotalAlloc-memStart.TotalAlloc)/1024)
}

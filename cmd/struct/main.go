package main

import (
	"fmt"
	"unsafe"
)

type Numbers struct {
	a int32
	b int64
}

func main() {
	// allocate a continuous block of memory for the struct
	var buf [unsafe.Sizeof(Numbers{})]byte
	ptr := unsafe.Pointer(&buf[0])

	// cast a pointer to the struct type
	s := (*Numbers)(ptr)

	// initialize fields
	s.a = 123
	s.b = 456
	fmt.Printf("a: %d | b: %d\n", s.a, s.b)

	// initialize fields with pointer arithmetic
	*(*int32)(ptr) = 789
	*(*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(int(0)))) = 101112
	fmt.Printf("a: %d | b: %d\n", s.a, s.b)
}

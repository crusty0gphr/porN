package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := [3]int{1, 2, 3}
	fmt.Printf("original: %v\n", data) // Output: [1 2 3]

	// pointer to the head
	ptr := unsafe.Pointer(&data[0])
	// end of the array
	end := uintptr(ptr) + uintptr(len(data)*int(unsafe.Sizeof(data[0])))

loop:
	*(*int)(ptr) *= 10
	// move to the next pointer
	ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(0))
	// continue
	if uintptr(ptr) < end {
		goto loop
	}

	fmt.Printf("modified: %v\n", data) // Output: [10 20 30]
}

package main

import (
	"errors"
	"fmt"
	"strings"
)

type CustomError struct {
	Msg string
	Err error
}

// Error method makes CustomError implement the error interface
func (e *CustomError) Error() string {
	return e.Msg
}

// Unwrap method allows for error chain traversal
func (e *CustomError) Unwrap() error {
	return e.Err
}

func main() {
	// Create a new CustomError instance
	// We're wrapping a simple error created with errors.New()
	// This creates an error chain: CustomError -> "wrapped error"
	errCustom := &CustomError{
		Msg: "This is a custom error",
		Err: errors.New("wrapped error"),
	}

	// Create another layer of error wrapping using fmt.Errorf
	// The %w verb indicates that we want to wrap errCustom
	// This creates: fmt error -> CustomError -> "wrapped error"
	err := fmt.Errorf("this is a standard error: %w", errCustom)

	// Unwrap one layer of the error chain
	// This removes the fmt error wrapper, exposing errCustom
	unwrappedErr := errors.Unwrap(err)

	var customErr *CustomError
	if errors.As(unwrappedErr, &customErr) {
		// This will print "This is a custom error" as a result
		fmt.Println("Handled custom error:", customErr.Msg)
		return
	}

	if strings.Contains(err.Error(), "custom") {
	}

	fmt.Println("Handled standard error:", unwrappedErr)
}

package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("something went wrong")
}

func doSomethingElse() error {
	return errors.New("something else went wrong")
}

func run() (err error) {
	defer func() {
		err = doSomethingElse()
	}()

	if err = doSomething(); err != nil {
		return
	}

	return
}

func main() {
	fmt.Println(run())
}

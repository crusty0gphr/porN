package main

import (
	"errors"
	"sync"
)

func doSomething() error {
	return errors.New("something went wrong")
}

func doSomethingElse() error {
	return errors.New("something else went wrong")
}

func run() error {
	var err error

	defer func() {
		err = doSomethingElse()
	}()

	if err = doSomething(); err != nil {
		return err
	}

	return nil
}

type SomeType struct {
}

func (s *SomeType) SomeMethod() {

}

func main() {
	//fmt.Println(run())
	var wg sync.WaitGroup

	var a *SomeType
	wg.Add(1)
	go func() {

		a.SomeMethod()
	}()
}

package main

import (
	"os"
	"runtime"
	"runtime/pprof"
)

type state string

const (
	stateInProgress = iota + 1
	statePending
	stateCompleted
)

//go:noinline
func makeArray() [4]state {
	return [4]state{
		stateInProgress: "in progress",
		statePending:    "pending",
		stateCompleted:  "completed",
	}
}

//go:noinline
func makeMap() map[int]state {
	return map[int]state{
		stateInProgress: "in progress",
		statePending:    "pending",
		stateCompleted:  "completed",
	}
}

func main() {
	arr := makeArray()
	m := makeMap()

	runtime.KeepAlive(arr)
	runtime.KeepAlive(m)

	f, _ := os.Create("heap.prof")
	pprof.WriteHeapProfile(f)
	f.Close()
}

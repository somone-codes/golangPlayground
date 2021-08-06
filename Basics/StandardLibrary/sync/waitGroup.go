package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func foo() {
	fmt.Println("Hello I am foo!")
	waitGroup.Done()
}

func main() {
	fmt.Println("Number of go routines is ", runtime.NumGoroutine())
	waitGroup.Add(1)
	go foo()
	fmt.Println("I am print after foo call")
	fmt.Println("Number of go routines is ", runtime.NumGoroutine())
	fmt.Println("Waiting for foo to complete.")
	waitGroup.Wait()
	fmt.Println("foo should be complete now to complete.")

}

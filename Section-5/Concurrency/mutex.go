package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count = 0
var wg sync.WaitGroup

func incrAndPrint() {
	tempCount := count
	tempCount += 1
	count = tempCount
	fmt.Println("Count value is ", count)
	fmt.Println("Number of go routines when ending a incrAndPrint ", runtime.NumGoroutine())
	wg.Done()
}

func main() {

	parallelCount := 1_000_000 // if this is really low then we don't see the race condition
	wg.Add(parallelCount)

	for i := 0; i < parallelCount; i++ {
		fmt.Println("Number of go routines when starting new incrAndPrint ", runtime.NumGoroutine())
		go incrAndPrint()
	}
	fmt.Println("Waiting for all incrAndPrint to complete ")
	wg.Wait()
	fmt.Println("final count value is ", count) // usually around 999993 for 1 mil
}

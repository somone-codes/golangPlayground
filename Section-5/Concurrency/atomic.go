package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var atomicCount int64 = 0
var _wg sync.WaitGroup

func atomicIncrAndPrint() {
	atomic.AddInt64(&atomicCount, 1)
	fmt.Println("Count value is ", atomicCount)
	fmt.Println("Number of go routines when ending a incrAndPrint ", runtime.NumGoroutine())
	_wg.Done()
}

func main() {

	parallelCount := 1_000_000 // if this is really low then we don't see the race condition
	_wg.Add(parallelCount)

	for i := 0; i < parallelCount; i++ {
		fmt.Println("Number of go routines when starting new incrAndPrint ", runtime.NumGoroutine())
		go atomicIncrAndPrint()
	}
	fmt.Println("Waiting for all incrAndPrint to complete ")
	_wg.Wait()
	fmt.Println("final count value is ", atomicCount) // usually around 999993 for 1 mil
}

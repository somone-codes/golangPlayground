package main

import (
	"fmt"
	"runtime"
)

func Foo() {
	fmt.Println("Hello I am foo!")
}

func main() {
	fmt.Println("Number of go routines is ", runtime.NumGoroutine())
	go Foo() // now Foo is executed as a routine concurrent to this main func
	fmt.Println("I am print after foo call")
	fmt.Println("Number of go routines is ", runtime.NumGoroutine())
	//notice how Foo has not executed, as we never gave a chance to exec / waiting for it to complete as main prog exited
}

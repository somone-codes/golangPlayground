package main

import "fmt"

func main() {
	f := func() {
		fmt.Println(" I am a function assigned to var")
	}

	f()

	//cant have same var f so assigned new
	f1 := func(x int) {
		fmt.Println(" I am a function assigned to var who accepts params")
	}
	f1(22)
}

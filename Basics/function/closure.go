package main

import "fmt"

func main() {

	{
		x := " I am a closure"
		fmt.Print(x)
	}

	// x is not accessible here!

	a := incrementer()
	b := incrementer()
	fmt.Println(a()) // 0
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(b()) // 0
	fmt.Println(b()) // 1

	//notice how the state of x is stored and maintained for subsequent usage.
	//also notice that b has its own state of x
}

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Function literal egs")
	func() {
		fmt.Print("I am a function literal/ anonymous function.")
	}()
}

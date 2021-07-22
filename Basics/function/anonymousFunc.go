package main

import "fmt"

func main() {

	func() {
		fmt.Println(" I am a anonymous func with no params")
	}()

	func() {
		fmt.Println(" I am a anonymous func with no params too")
	}()

	func(x int) {
		fmt.Printf(" I am a anonymous func with params , %v", x)
	}(22)
}

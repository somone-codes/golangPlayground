package main

import "fmt"

func main() {
	//for loop and variations
	fmt.Println("\nFor loop example")
	i := 1
	fmt.Println("simple For loop")
	for {
		if i > 2 {
			break
		}
		fmt.Print(i)
		i++
	}
	fmt.Println("\nFor loop with init, condition and incr")
	//notice doubt short hand op as I was initialized before as well
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	i = 0
	fmt.Println("\nFor loop with condition")
	for i < 3 {
		fmt.Print(i)
		i++
	}

	// for loop through collection like slice/array..
	fmt.Println()
	array := [...]string{" Pen", "Teller"}
	for length, value := range array { // value is optional
		fmt.Println(length, value)
	}

}

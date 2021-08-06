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

	//for loop channel
	ch := make(chan int)
	go func(writeChannel chan<- int) {
		for i := 0; i < 100; i++ {
			writeChannel <- i
		}
		//can close only write and read-write channel i.e read only channels can't be closed!
		close(writeChannel) // remove this and see what happens?
	}(ch)

	for data := range ch { // keeps reading till channel is closed.
		fmt.Println(data)
	}

}

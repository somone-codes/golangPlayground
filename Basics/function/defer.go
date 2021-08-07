package main

import "fmt"

func main() {
	fmt.Println(plus1(1))
	fmt.Println(plus1(120))
}

func plus1(i int) (j int) { // notice the return variable
	defer func() { j++ }() // exec after the return , incr j and this will be returned
	j = i
	return j // setting j = 1
}

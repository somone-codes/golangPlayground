package main

import "fmt"

func main() {
	slice := []string{"I am", "a slice"}
	fmt.Println(slice)
	fmt.Printf("type of slice is %T", slice)

	// creating slice using make func

	var makeSliceWithCap []string
	makeSliceWithCap = make([]string, 1, 2)
	makeSliceWithCap[0] = "I am a slice too!"
	makeSliceWithCap[1] = "I am a slice element 2!"
	//makeSliceWithCap[2] = "I am a slice element 3!" // index out of bound as its exceeding cap 2
	fmt.Println(" \n", makeSliceWithCap)
	fmt.Println(" len of slice is", len(makeSliceWithCap))
	fmt.Println(" cap of slice is", cap(makeSliceWithCap))

	var makeSliceWithoutCap []string
	makeSliceWithoutCap = make([]string, 1)
	fmt.Println(" \n", makeSliceWithoutCap)
}

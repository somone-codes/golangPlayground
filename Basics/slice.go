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
	// makeSliceWithCap[1] = "I am a slice element 2!" // index out of bound as its exceeding cap 2 but you can append 1 more
	//makeSliceWithCap[2] = "I am a slice element 3!"
	fmt.Println(" \n", makeSliceWithCap)
	fmt.Println(" len of slice is", len(makeSliceWithCap))
	fmt.Println(" cap of slice is", cap(makeSliceWithCap))

	var makeSliceWithoutCap []string
	makeSliceWithoutCap = make([]string, 1)
	fmt.Println(" Default value is nil", makeSliceWithoutCap)
	fmt.Println(" len of makeSliceWithoutCap is", len(makeSliceWithoutCap))
	fmt.Println(" cap of makeSliceWithoutCap is", cap(makeSliceWithoutCap))

	// creating slice from a array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x, start and end index in between : is optional
	fmt.Println(" slice of array s is", s)
}

package main

import "fmt"

func main() {
	var array [2]int
	fmt.Println(array)
	array[0] = 20
	fmt.Println(array)

	array2 := [...]string{" Pen", "Teller"} // auto determine length
	fmt.Println(array2)
	fmt.Printf("type of array2 is %T", array2)
	fmt.Println(" len of array2 is", len(array2))
}

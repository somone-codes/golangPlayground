package main

import "fmt"

func main() {
	//short declaration operator
	shortOp := "I am short op" // type is inferred and changing value of this var in future must be string only for string
	fmt.Println(shortOp)
	fmt.Printf("Type of variable shortOp is  %T ", shortOp) // using printf ->print with format
}

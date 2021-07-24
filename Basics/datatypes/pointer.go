package main

import "fmt"

func main() {
	x := 10
	b := &x                //obtaining address of x
	var pointer **int = &b // declaring pointer to pointer

	fmt.Printf("\ntype of b is %T\n", b)
	fmt.Println("Address of x in b is ", b) // *int -> implies pointer of type int
	fmt.Println("Address of x is ", &x)
	fmt.Println("Address of b is ", &b)
	fmt.Println("Obtaining value of x though address in b by dereference is ", *b)
	fmt.Println("value in x is ", *&x)
	fmt.Println("Obtaining value of x though pointer which holder addresss of b which holds address of x is ",
		**pointer)

	*b = 11 // changing value of x though b
	fmt.Println("Value in x now is ", x)
}

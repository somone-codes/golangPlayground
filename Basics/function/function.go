package main

import "fmt"

func main() {
	combineParamsOfSameType("s1", "s2", 1)
	withReturn("s1") // need not accept the return
	s1, i1 := withMultipleReturns()
	fmt.Println(s1, i1)
	variadicParameter(1, 2, 3, 4)

	defer withParams("\nI am called ", 2) //executed last after all functions/lines below
	withParams("I am called ", 1)         // with defer this line is executed before the func with deffer and on completion
	// or failure defer func is called

	// function that returns a function
	x := functionReturnsFunc()
	fmt.Printf("\n type of x is %T", x)
	x(10)
	functionReturnsFunc()(1)
}

//pass by value but if slice or array or etc is used then its reference/pointer is passed so changes will be seen by
//caller
//need not call all defined func
func withParams(s1 string, i1 int) {
	fmt.Println(s1, i1)
}

func combineParamsOfSameType(s1, s2 string, i1 int) { // can combine vars of same type together optionally and put type
	fmt.Println(s1, i1)
	//need not use s2
}

//variadic param i is a slice of int (zero or more)
//variadic param has to be the last param and can have other params before it
func variadicParameter(i ...int) {

	fmt.Printf("\nType of i is %T, value is %v\n", i, i)
}

func withReturn(s1 string) string {
	return "I am a string"
}

func withMultipleReturns() (string, int) {
	return "I am a string", 1
}

//this function returns a func that take int, prints it and returns a bool
func functionReturnsFunc() func(i int) bool {
	return func(i int) bool {
		fmt.Println("\n i is ", i)
		return true
	}
}

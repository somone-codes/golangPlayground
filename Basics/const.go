package main

import "fmt"

const constant = "I am a const!"

const (
	const1         = 42
	const2         = "I am a constant too!" // untyped constant
	const3 float32 = 23.44                  // typed constant
)

func main() {

	const constantDeclaredButUnused = 22 // we can have constants that are declared but unused but not vars

	fmt.Println(constant)
}

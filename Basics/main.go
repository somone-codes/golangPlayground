package main

import "fmt"

var variable = 22 // vars here can go unused wont throw err only block ones throw err

// x:=11 this is not allow, only allowed in block level

const constant = "I am a const!"

const (
	const1         = 42
	const2         = "I am a constant too!" // untyped constant
	const3 float32 = 23.44                  // typed constant
)

var (
	multiVar1      = "I am a multi Var!"
	multiVar2 int8 = 22
)

const (
	u         = iota      // u == 0     (untyped integer constant)  iota = 0
	v float64 = iota * 42 // v == 42.0  (float64 constant)			 iota = 1
	w         = iota      // w == 84    (untyped integer constant)	 iota = 2
	a         = 1 << iota // a == 8							     iota = 3
	b                     // b == 16		1 << a					 iota = 4
	c                     // c == 32		1 << b					 iota = 5
	d = iota              // d == 6								 iota = 6
	e = 23                // e = 23								 iota = 7
	f                     // f = 23								 iota = 8
	g = iota              // g = 9									 iota = 9
)

func main() { // this bracket has to open here only else compilation fails
	var variableWithTypeDef int // non usage of declared variable throws err

	//short declaration operator
	shortOp := "I am short op" // type is inferred and changing value of this var in future must be string only for string
	fmt.Println(shortOp)

	fmt.Println("We can use vars outside of function but no shorthand ops ", variable)

	fmt.Println("Variable with typedef have deefault values equal to 0 based on type store in them ",
		variableWithTypeDef)

	fmt.Printf("Type of variable shortOp is  %T ", shortOp) // using printf ->print with format

	rawString := `I am a raw string.
						"This is what I can doooo"`

	fmt.Print(rawString)

	type customType int // defined my own type based on int (sorta like extends int)
	var customTypeVariable customType = 10
	fmt.Printf("\nType of customTypeVar is %T and value is %v", customTypeVariable, customTypeVariable)

	var typeConversion int = int(customTypeVariable) // since underlying type of customType is int we can do this
	fmt.Print("Type conversion ", typeConversion)

	fmt.Println(constant)

	const constantDeclaredButUnused = 22 // we can have constants that are declared but unused but not vars

	fmt.Println(u, v, w, a, b, c, d, e, f, g)

	const (
		b  = iota + 1
		kb = 1 << (iota * 10) // left shift 1 by 10 times i.e  00000000 00000001 -> 00000100 00000000 i.e 1 becomes 1024
		mb
		gb
	)

	fmt.Printf("%d\t\t%b\n", b, b)
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	const x = 10
	fmt.Printf("%d\t%b\n", x, x)

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

	// if statement
	if something := 22; something == 22 {
		fmt.Print("something here is local var, and this if can have assignment + condition")
	}

}

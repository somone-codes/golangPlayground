package main

import "fmt"

var variable = 22

func main() { // this bracket has to open here only else compilation fails
	var variableWithTypeDef int // non usage of declared variable throws err
	print("hello")

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

}

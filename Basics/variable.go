package main

import "fmt"

var variable = 22 // vars here can go unused wont throw err only block ones throw err

// x:=11 this is not allow, only allowed in block level

var (
	multiVar1      = "I am a multi Var!"
	multiVar2 int8 = 22
)

func main() {

	fmt.Println("We can use vars outside of function but no shorthand ops ", variable)

	var variableWithTypeDef int // non usage of declared variable throws err

	fmt.Println("Variable with typedef have default values equal to 0 based on type store in them ",
		variableWithTypeDef)

}

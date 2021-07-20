package main

var variable = 22

func main() { // this bracket has to open here only else compilation fails
	var variableWithTypeDef int // non usage of declared variable throws err
	print("hello")

	//short declaration operator
	shortOp := "I am short op" // type is inferred and changing value of this var in future must be string only for string
	println(shortOp)

	println("We can use vars outside of function but no shorthand ops ", variable)

	println("Variable with typedef have deefault values equal to 0 based on type store in them ",
		variableWithTypeDef)
}

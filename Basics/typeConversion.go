package main

func main() {
	type customType int // defined my own type based on int (sorta like extends int)
	var customTypeVariable customType = 10
	fmt.Printf("\nType of customTypeVar is %T and value is %v", customTypeVariable, customTypeVariable)

	var typeConversion int = int(customTypeVariable) // since underlying type of customType is int we can do this
	fmt.Print("Type conversion ", typeConversion)

}

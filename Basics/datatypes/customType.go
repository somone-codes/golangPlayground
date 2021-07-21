package main

import "fmt"

func main() {
	type customType int // defined my own type based on int (sorta like extends int)
	var customTypeVariable customType = 10
	fmt.Printf("\nType of customTypeVar is %T and value is %v", customTypeVariable, customTypeVariable)

}

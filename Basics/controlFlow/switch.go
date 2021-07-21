package main

import "fmt"

func main() {
	// switch statement
	i := 20
	//switch without expression

	switch {
	case 1 == 2:
		fmt.Print(" false")
	case 1 == 1:
		fmt.Println("\ntrue")
		fmt.Println("1==1")
		fallthrough
	case i == 2:
		fmt.Print(" false")
	default:
		fmt.Print("default case")
	}

	//switch with initialization

	switch some := 12; { // this init can as be return value from func i.e some:=func()
	case 1 == 1:
		fmt.Print(some)
	}

	//switch with checks
	switch i {
	case 2:
		fmt.Print(" i ==2")
	default:
		fmt.Printf("i is %v\n", i)
	}

	// type switch tbd
}

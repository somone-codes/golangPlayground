package main

import "fmt"

func main() {
	fmt.Print(opOnEvenNumber(sum, 1, 3, 2, 5, 6, 8, 8))
}

func sum(i ...int) int {
	sum := 0
	for _, i2 := range i {
		sum += i2
	}
	return sum
}

//this function accepts CALLBACK and variadic int and applies that function to all even numbers in variadic
//int parameter
func opOnEvenNumber(f func(num ...int) int, i ...int) int {
	var intSlice []int
	fmt.Println(i)
	for _, i2 := range i {
		if i2%2 == 0 {
			intSlice = append(intSlice, i2)
		}
	}
	return f(intSlice...)
}

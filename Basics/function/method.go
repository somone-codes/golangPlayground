package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int8
	sex   string
	id    int
} // we need not use struct we can use any type, even custom type in for methods

//getName is  a method
//p is receiver
//whoever is the receiver implies this method is part of that type
func (p person) getName() string {
	return fmt.Sprint(p.fname, p.lname)
}

func main() {
	p1 := person{
		fname: "Someone",
		lname: "Learns",
		age:   22,
		sex:   "Secret",
		id:    0,
	}
	fmt.Printf(p1.getName())
}

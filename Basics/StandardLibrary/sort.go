package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int8
	sex  string
	id   int
}

func (p person) String() string {
	return fmt.Sprintf("{Name: %s, id: %d , age: %d },", p.name, p.age, p.id)
} // overriding default string func, this is like toString in java and called when we string access value of person type

type peopleIdSort []person

func (p peopleIdSort) Len() int           { return len(p) }
func (p peopleIdSort) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p peopleIdSort) Less(i, j int) bool { return p[i].id < p[j].id }

func main() {

	// int sort
	var intSlice = []int{44, 66, 23, 11, 1, 9, 10, 99}

	fmt.Println(" intSlice before sort ", intSlice)
	sort.Ints(intSlice)
	fmt.Println(" intSlice after sort ", intSlice)

	//similarly for string and others we have like sort.String/floats etc

	//custom sort
	p1 := person{
		name: "someone",
		age:  22,
		sex:  "Secret",
		id:   0,
	}

	p2 := person{
		name: "someOtherOne",
		age:  25,
		id:   3,
	}
	p3 := person{
		name: "someOtherTwo",
		age:  98,
		id:   1,
	}

	people := []person{p2, p1, p3}
	fmt.Println("People before sort by ID is ", people)
	sort.Sort(peopleIdSort(people))
	fmt.Println("People after sort by ID is ", people)
}

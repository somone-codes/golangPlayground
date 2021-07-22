package main

import "fmt"

//person struct
//there is already 1 person struct in method.go and both are in same package function hence making it _person else error
type _person struct {
	fname          string
	lname          string
	age            int8
	sex            string
	id             int64
	height, weight int32
}

//person getId
func (pep _person) getId() string {
	return fmt.Sprint(pep.id)
}

//student struct
type student struct {
	_person
	id int64
}

//student getId
func (stud student) getId() string {
	return fmt.Sprint(stud.id)
}

// interface, whoever implements getId func with return type string is of type studentPerson
// here both student and person have a method getId
type studentPerson interface {
	getId() string
}

// polymorphism accepts both student and person type
// anyone value is of type studentPerson can call this method
func getInfo(studPep studentPerson) string {

	//switch case with type
	switch studPep.(type) {
	case _person:
		// studPep.(_person) this is INSERTION, similar to CONVERSION
		return fmt.Sprint(studPep.(_person).fname, " ", studPep.(_person).lname, " ", studPep.(_person).id)
	case student:
		return fmt.Sprint(studPep.(student).fname, " ", studPep.(student).lname, " ", studPep.(student).id)
	default:
		return fmt.Sprint("Oops no matching type found!")
	}
}

func main() {

	// anything after :=  is called a value of the type
	//p1 value is of type person and studentPerson
	p1 := _person{
		fname: "Someone",
		lname: "Learns",
		age:   22,
		sex:   "Secret",
		id:    0,
	}

	s1 := student{
		_person: p1,
		id:      32,
	} // s1 value is of type student and studentPerson with embedded type person

	fmt.Printf("Person p1 id is %v\n\n", p1.getId())
	fmt.Printf("Student s1 id is %v\n", s1.getId())

	fmt.Printf("Person p1 complete info is %v\n", getInfo(p1))
	fmt.Printf("Student s1 complete info is %v\n\n", getInfo(s1))
}

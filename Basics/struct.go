package main

import "fmt"

func main() {
	type person struct {
		name           string
		age            int8
		sex            string
		id             int //we can also have functions like F func() and pointers *P .Type
		height, weight int32
	}

	p1 := person{
		name: "someone",
		age:  22,
		sex:  "Secret",
		id:   0,
	} // p1 is called a value of type person and its not referred as object in go

	fmt.Println(p1)
	fmt.Println(p1.age)

	// Embedded struct

	type student struct {
		person // anonymous field since type is no specified specifically
		id     int64
	}

	s1 := student{
		person: p1,
		id:     32,
	}

	s2 := student{
		person: person{
			name: "someOtherOne",
			age:  25,
			id:   3, // purposely skipping sex and it doesnt throw err
		},
		id: 44,
	}

	fmt.Println(s1)
	fmt.Println(s1.age)       //
	fmt.Println(s1.person.id) // since there is name collision for id FILED going to INNER type and getting it

	fmt.Println(s2)

	// anonymous struct
	anonymousStruct := struct {
		name           string
		age            int8
		sex            string
		id             int //we can also have functions like F func() and pointers *P .Type
		height, weight int32
	}{
		name: "anonymous",
		age:  29,
	}

	fmt.Println(anonymousStruct)
}

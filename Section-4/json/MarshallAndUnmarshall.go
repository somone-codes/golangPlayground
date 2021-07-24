package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string // notice the fields start with caps unlike the one in struct,else marshall return empty
	Age  int8
	Sex  string
	Id   int
}

func marshallAndPrint(data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n type of marhsall response is %T \n", jsonData)
		fmt.Println(string(jsonData)) // converting to string as response is []byte/ []uint8, byte is alias for uint8
	}
}
func main() {
	p1 := person{
		Name: "someone",
		Age:  22,
		Sex:  "Secret",
		Id:   0,
	}

	marshallAndPrint(p1) // json

	p2 := person{
		Name: "someOtherOne",
		Age:  25,
		Id:   3, // purposely skipping sex and it doesnt throw err
	}

	people := []person{p1, p2}

	marshallAndPrint(people) // json Array

}

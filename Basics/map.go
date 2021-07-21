package main

import "fmt"

func main() {
	nameIdMap := map[string]int{
		"Someone":   1,
		"SomeOther": 2,
	}

	fmt.Println(nameIdMap["Someone"])
	fmt.Println(nameIdMap["I dont exist"]) // for non existing key it give 0

	//to check if key exists in map (LOOK UP)
	val, ok := nameIdMap["I dont exist"] // comma ok idiom
	fmt.Println(val)
	fmt.Println(ok) // false since element doesnt exists

	//same as above written in if
	if value, okay := nameIdMap["SomeOther"]; okay {
		fmt.Println(value)
	}

	// adding element
	nameIdMap["NewElement"] = 22

	//deleting
	delete(nameIdMap, "Someone") // if key doesnt exists behaviour is same as lookup above
	fmt.Println(nameIdMap)
}

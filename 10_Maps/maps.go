package main

import (
	"fmt"
	"maps"
)

func main() {
	// Creating map
	m := make(map[string]string)

	// Selecting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// Get an element
	fmt.Println(m["name"], m["area"])

	m1 := make(map[string]int)

	m1["age"] = 30
	m1["price"]  =50

	fmt.Println(m1["age"])
	fmt.Println(len(m1))

	delete(m, "area")
	fmt.Println(m)


	m2 := map[string]int {"price":40, "phones":3}

	val, ok := m2["phones"]
	fmt.Println(val)
	if ok{
		fmt.Println("all ok")
	} else{
		fmt.Println("Not Ok")
	}

	m3 := map[string]int {"price":40, "phones":3}

	fmt.Println(maps.Equal(m2, m3))

}

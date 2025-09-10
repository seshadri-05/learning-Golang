package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating maps

	// m := make(map[string]string)

	// setting an element
	// m["name"] = "Golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])

	// if key does not exist in the map then it returns zero value
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m, "price")
	// fmt.Println(m)

	// another way of creating map
	// a := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(a)

	/* a := map[string]int{"price": 40, "phones": 3}

	_, ok := a["phone"]
	if ok {
		fmt.Println("all ok")

	} else {
		fmt.Print("not ok")
	} */

	m1 := map[string]int{"price": 40, "phones": 5}
	m2 := map[string]int{"price": 40, "phones": 5}

	fmt.Println(maps.Equal(m1, m2))
}

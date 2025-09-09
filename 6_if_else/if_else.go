package main

import "fmt"

func main() {

	/* age := 21

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {

		fmt.Println("Person is not an adult")
	}
	*/

	age := 21

	if age >= 18 {
		fmt.Println("Person is an adult")

	} else if age >= 12 {

		fmt.Println("Person is a teenager")

	} else {
		fmt.Println("Person is a kid")
	}

	// OR operator
	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}

	// we can declare a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("Peron is an adult", age)

	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}
}

// go doesnt have ternary, will need to use normal if else

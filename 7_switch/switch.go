package main

import (
	"fmt"
)

func main() {

	//simple switch
	/* i := 5

	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("Other")
	} */

	// multiple condition switch

	/* switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")

	default:
		fmt.Println("Weekday")

	} */

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an Integer")

		case string:
			fmt.Println("Its a string")

		case bool:
			fmt.Println("Its boolean")

		default:
			fmt.Println("Other", t)

		}
	}

	whoAmI(90.06)
}

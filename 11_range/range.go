package main

import "fmt"

func main() {
	// nums := []int{6, 7, 8}

	// sum := 0

	//for i, num := range nums {
	// sum = sum + num
	//	fmt.Println(num, i)
	//}

	/* m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	} */

	// unicode code
	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
}

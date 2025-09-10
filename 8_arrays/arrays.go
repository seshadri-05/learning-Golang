package main

import "fmt"

func main() {
	/* var nums [4]int
	nums[0] = 1

	fmt.Println(len(nums)) // array length
	fmt.Println(nums[0])
	fmt.Println(nums) */

	// to declare it in a single line
	nums := [4]int{1, 2, 3}

	fmt.Println(nums)

	// 2D Arrays
	vals := [2][2]int{{3, 4}, {3, 4}}

	fmt.Println(vals)

}

// int -> 0, string -> "", bool -> False

package main

import "fmt"

/* func changeNum(num int) {

	num = 5
	fmt.Println("In changeNum", num)
}
*/

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)

}

func main() {

	num := 1
	// changeNum(num)
	changeNum(&num)

	//fmt.Println("Memory Adress", &num)
	fmt.Println("After changeNum in main", num)

}

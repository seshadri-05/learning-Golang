package main

/* func add(a int, b int) int {
	return a + b
} */
/*
func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}
*/

// func processIt(fn func(a int) int) {
//	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	/* result := add(5, 3)
	fmt.Println(result) */

	/* fn := func(a int) int {
		return 2
	}
	processIt(fn) */


	

	fn := processIt()
	fn(6)

}

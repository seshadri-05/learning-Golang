package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialized slice is nil
	/* var nums []int

	fmt.Println(len(nums)) */

	// var nums = make([]int, 2, 5)
	// var nums = make([]int, 0, 5)
	// nums := []int{}
	// capacity -> maximum numbers of elements can fit
	/* nums[0] = 45
	nums[1] = 56
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	fmt.Println(nums)
	fmt.Println(cap(nums))


	*/

	/* var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))

	// copy function
	copy(nums2, nums)
	fmt.Println(nums, nums2)
	*/

	// slice operator

	var nums = []int{1, 2, 3, 4, 5}

	fmt.Println(nums[1:])

	// slice package
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2))

	

	// 2D slices
	var nums3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums3)

}
